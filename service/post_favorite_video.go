package service

import (
	"errors"
	"tiktok/cache"
	"tiktok/models"
)

const(
	PLUS = 1	//点赞
	MINUS = 2	//取消
)

func PostFavorite(userId int32,videoId int32,actionType int32) error{
	return NewFavoriteFlow(userId,videoId,actionType).Do()
}

func NewFavoriteFlow(userId int32,videoId int32,actionType int32) *FavoriteFlow{
	return &FavoriteFlow{
		UserId:userId,
		VideoId: videoId,
		ActionType: actionType,
	}
}

type FavoriteFlow struct {
	UserId int32
	VideoId int32
	ActionType int32
}


func (f *FavoriteFlow) Do() error{
	if err := f.Check();err != nil{
		return err
	}
	if err := f.prepareData();err != nil{
		return err
	}

	return nil
}

// Check 参数检查
func (f *FavoriteFlow) Check() error{
	if !models.NewUserInfoDAO().IsUserExitsById(f.UserId){
		return errors.New("user not found")
	}

	if !models.NewVideoDAO().IsVideoExitsByVideoId(f.VideoId){
		return errors.New("video not found")
	}

	if f.ActionType != PLUS && f.ActionType != MINUS{
		return errors.New("actionType err")
	}

	return nil
}

//数据处理
func (f *FavoriteFlow) prepareData() error{

	switch f.ActionType {
	case PLUS:
		if err := f.PlusOperation();err != nil{
			return err
		}
	case MINUS:
		if err := f.MinusOperation();err != nil{
			return err
		}
	default:
		return errors.New("actionType err")
	}

	return nil
}

// PlusOperation 增加赞操作
func (f *FavoriteFlow) PlusOperation() error{
	redisOp := cache.NewRedisOperation()
	dao := models.NewVideoDAO()

	//redis操作
	if err := redisOp.AddFavorite(f.UserId,f.VideoId);err != nil{
		return err
	}

	//todo 可以只使用redis？？

	//mysql操作

	//视频增加赞
	err := dao.AddVideoFavorite(f.VideoId,f.UserId)
	if err != nil{
		return err
	}

	return nil
}

// MinusOperation 减少赞操作
func (f *FavoriteFlow) MinusOperation() error{
	redisOp := cache.NewRedisOperation()
	dao := models.NewVideoDAO()

	//redis操作
	if err := redisOp.CancelFavorite(f.UserId,f.VideoId);err != nil{
		return err
	}

	//mysql操作
	err := dao.CancelVideoFavorite(f.VideoId,f.UserId)
	if err != nil{
		return err
	}

	return nil
}