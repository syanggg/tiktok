package service

import (
	"errors"
	"tiktok/models"
)

type FollowerListResponse struct {
	FollowerList []*models.UserInfo `json:"user_list"`
}

func QueryFollowerList(userId int32) (*FollowerListResponse,error){
	return NewFollowerListFlow(userId).Do()
}

func NewFollowerListFlow(userId int32) *FollowerListFlow{
	return &FollowerListFlow{
		UserId: userId,
	}
}

type FollowerListFlow struct {
	UserId int32

	*FollowerListResponse
	followerList []*models.UserInfo
}

func (f *FollowerListFlow) Do() (*FollowerListResponse,error){
	if err := f.Check();err != nil{
		return nil,err
	}
	if err := f.prepareData();err != nil{
		return nil,err
	}

	f.packResponse()

	return f.FollowerListResponse,nil
}


func (f *FollowerListFlow) Check() error{
	if !models.NewUserInfoDAO().IsUserExitsById(f.UserId){
		return errors.New("user not found")
	}

	return nil
}

func (f *FollowerListFlow) prepareData() error{
	dao := models.NewUserInfoDAO()

	//获取粉丝列表
	list,err := dao.GetFollowerListById(f.UserId)
	if err != nil{
		return err
	}

	f.followerList = list

	if err := fillUserFollow(f.UserId,list);err != nil{
		return err
	}

	return nil
}

func (f *FollowerListFlow) packResponse(){
	f.FollowerListResponse = &FollowerListResponse{
		FollowerList: f.followerList,
	}
}
