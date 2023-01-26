package service

import (
	"errors"
	"tiktok/cache"
	"tiktok/models"
)

const (
	FOLLOW = 1
	UNFOLLOW = 2
)

func PostFollow(userId int32,followId int32,actionType int32) error{
	return NewFollowFlow(userId,followId,actionType).Do()
}

func NewFollowFlow(userId int32,followId int32,actionType int32) *FollowFlow{
	return &FollowFlow{
		UserId: userId,
		FollowId: followId,
		ActionType: actionType,
	}
}

type FollowFlow struct {
	UserId int32
	FollowId int32
	ActionType int32
}

func (f *FollowFlow) Do() error{
	if err := f.Check();err != nil{
		return err
	}

	if err := f.prepareData();err != nil{
		return err
	}

	return nil
}

func (f *FollowFlow) Check() error{
	dao := models.NewUserInfoDAO()

	if !dao.IsUserExitsById(f.UserId){
		return errors.New("user not found")
	}

	if !dao.IsUserExitsById(f.FollowId){
		return errors.New("follow user not found")
	}

	if f.ActionType != FOLLOW && f.ActionType != UNFOLLOW{
		return errors.New("actionType err")
	}

	return nil
}

func (f *FollowFlow) prepareData() error{
	switch f.ActionType {
	case FOLLOW:
		if err := f.FollowOperation();err != nil{
			return err
		}
	case UNFOLLOW:
		if err := f.UnfollowOperation();err != nil{
			return err
		}
	default:
		return errors.New("actionType err")
	}

	return nil
}

// FollowOperation 关注操作
func (f *FollowFlow) FollowOperation() error{
	redisOp := cache.NewRedisOperation()
	dao := models.NewUserInfoDAO()

	//redis操作
	err := redisOp.AddFollow(f.UserId,f.FollowId)
	if err != nil{
		return err
	}

	//mysql操作
	if err := dao.AddUserFollow(f.UserId,f.FollowId);err != nil{
		return err
	}

	return nil
}

// UnfollowOperation 取消关注操作
func (f *FollowFlow) UnfollowOperation() error{
	redisOp := cache.NewRedisOperation()
	dao := models.NewUserInfoDAO()

	//redis操作
	err := redisOp.CanCelFollow(f.UserId,f.FollowId)
	if err != nil{
		return err
	}

	//mysql操作
	if err := dao.CancelUserFollow(f.UserId,f.FollowId);err != nil{
		return err
	}

	return nil
}