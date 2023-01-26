package service

import (
	"errors"
	"tiktok/cache"
	"tiktok/models"
)

type FollowListResponse struct {
	 FollowList []*models.UserInfo `json:"user_list"`
}

func QueryFollowList(userId int32) (*FollowListResponse,error){
	return NewFollowListFlow(userId).Do()
}

func NewFollowListFlow(userId int32) *FollowListFlow{
	return &FollowListFlow{
		UserId: userId,
	}
}

type FollowListFlow struct {
	UserId int32

	*FollowListResponse
	followList []*models.UserInfo
}

func (f *FollowListFlow) Do() (*FollowListResponse,error){
	if err := f.Check();err != nil{
		return nil,err
	}

	if err := f.prepareData();err != nil{
		return nil,err
	}

	f.packResponse()

	return f.FollowListResponse,nil
}

func (f *FollowListFlow) Check() error{
	if !models.NewUserInfoDAO().IsUserExitsById(f.UserId){
		return errors.New("user not found")
	}

	return nil
}

func (f *FollowListFlow) prepareData() error{
	dao := models.NewUserInfoDAO()

	//获取关注列表
	list,err := dao.GetFollowListById(f.UserId)
	if err != nil{
		return err
	}

	//填充用户关注状态
	if err := fillUserFollow(f.UserId,list);err != nil{
		return err
	}

	f.followList = list

	return nil
}

func (f *FollowListFlow) packResponse(){
	f.FollowListResponse = &FollowListResponse{
		FollowList: f.followList,
	}
}

//填充关注状态
func fillUserFollow(userId int32,users []*models.UserInfo) error{
	redis := cache.NewRedisOperation()

	for _,user := range users{
		isFollow,err := redis.GetFollowState(userId,user.Id)
		if err != nil{
			return err
		}

		user.IsFollow = isFollow
	}


	return nil
}