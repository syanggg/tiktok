package service

import "tiktok/models"

type UserInfoResponse struct {
	*models.UserInfo `json:"user"`
}

func QueryUserInfo(id int32) (*UserInfoResponse,error){
	return NewUserInfoFlow(id).Do()
}

func NewUserInfoFlow(id int32) *UserInfoFlow{
	return &UserInfoFlow{
		userId: id,
	}
}

type UserInfoFlow struct {
	userId int32
	*UserInfoResponse

	*models.UserInfo
}

func (u *UserInfoFlow) Do() (*UserInfoResponse,error){
	//models层数据库查询
	err := u.prepareData()
	if err != nil{
		return nil,err
	}
	//打包response数据
	u.packResponse()

	return u.UserInfoResponse,nil

}

func (u *UserInfoFlow) prepareData() error{

	dao := models.NewUserInfoDAO()
	info,err := dao.GetUserInfoById(u.userId)
	if err != nil{
		return err
	}

	u.UserInfo = info

	return nil
}

func (u *UserInfoFlow) packResponse(){
	u.UserInfoResponse = &UserInfoResponse{
		UserInfo:u.UserInfo,
	}
}