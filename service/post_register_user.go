package service

import (
	"errors"
	"tiktok/middleware"
	"tiktok/models"
)

const(
	MaxUserNameSize = 100
	MinPasswordSize = 6
	MaxPasswordSize = 20
)

func PostRegister(username ,password string) (*LoginInfoResponse,error) {
	return NewPostRegisterFlow(username,password).Do()
}

func NewPostRegisterFlow(username,password string) *PostRegisterFlow{
	return &PostRegisterFlow{
		username: username,
		password: password,
	}
}

type PostRegisterFlow struct {
	username string
	password string

	data *LoginInfoResponse
	userid int32
	token string
}

func (p *PostRegisterFlow) Do() (*LoginInfoResponse,error){
	//对参数进行合法验证
	err := p.Check()
	if err != nil{
		 return nil,err
	}
	//更新数据库
	if err := p.prepareData();err != nil{
		return nil,err
	}

	//打包response
	p.PackResponse()

	return p.data, nil
}

// Check 对参数进行合法验证
func (p *PostRegisterFlow) Check() error{
	if p.username == ""{
		return errors.New("username is a empty string")
	}
	if len(p.username) > MaxUserNameSize{
		return errors.New("MaxUserNameSize err")
	}
	if p.password == ""{
		return errors.New("password is a empty string")
	}

	return nil
}

// 数据处理
func (p *PostRegisterFlow) prepareData() error{
	login := models.NewUserLogin(p.username,p.password)
	info := models.NewUserInfo(p.username,login)

	//判断username是否存在
	loginDao := models.NewUserLoginDAO()
	if loginDao.IsUserExistByUsername(p.username){
		return errors.New("username is Already exist")
	}

	//添加 info表和login表一对一关系 login表存在外键 添加info表自动添加login表
	infoDao := models.NewUserInfoDAO()
	if err := infoDao.AddUserInfo(info);err != nil{
		return err
	}

	p.userid = info.Id
	//颁发token
	token,err := middleware.ReleaseToken(login)
	if err != nil{
		return err
	}
	p.token = token

	return nil
}

// PackResponse 打包loginResponse
func (p *PostRegisterFlow) PackResponse(){
	 p.data = &LoginInfoResponse{
		 UserId: p.userid,
		 Token: p.token,
	 }
}