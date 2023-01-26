package service

import (
	"errors"
	"tiktok/middleware"
	"tiktok/models"
)

type LoginInfoResponse struct{
	UserId int32 `json:"user_id"`
	Token string `json:"token"`
}

// QueryLogin 开始处理login服务
func QueryLogin(username,password string) (*LoginInfoResponse,error){
	return NewQueryLoginFlow(username,password).Do()
}

func NewQueryLoginFlow(username,password string) *QueryLoginFlow{
	return &QueryLoginFlow{
		username: username,
		password: password,
	}
}

type QueryLoginFlow struct {
	username string
	password string

	data *LoginInfoResponse
	userId int32
	token string
}


func (q *QueryLoginFlow) Do() (*LoginInfoResponse,error){
	//对参数进行合法检验
	err := q.Check()
	if err != nil{
		return nil,err
	}
	//数据库查询当前登录用户
	if err := q.prepareData();err != nil{
		return nil,err
	}

	//打包response
	q.packResponse()

	return q.data,nil
}

func (q *QueryLoginFlow) Check() error{
	if q.username == ""{
		return errors.New("username is a empty string")
	}
	if !models.NewUserLoginDAO().IsUserExistByUsername(q.username){
		return errors.New("user not found")
	}

	if len(q.username) > MaxUserNameSize{
		return errors.New("username MaxUserNameSize err")
	}
	if q.password == ""{
		return errors.New("password is a empty string")
	}

	return nil
}

//数据操作 查询当前登录用户
func (q *QueryLoginFlow) prepareData() error{

	//查询登录用户
	dao := models.NewUserLoginDAO()
	login,err := dao.QueryUserLogin(q.username,q.password)
	if err != nil{
		return err
	}

	q.userId = login.UserInfoId

	//发放token
	token,err := middleware.ReleaseToken(login)
	if err != nil{
		return err
	}
	q.token = token

	return nil
}

func (q *QueryLoginFlow) packResponse(){
	q.data = &LoginInfoResponse{
		UserId: q.userId,
		Token: q.token,
	}
}