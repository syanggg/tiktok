package models

import (
	"log"
	"sync"
)

type UserLogin struct {
	Id         int32  `gorm:"primary_key"`
	Username   string `gorm:"primary_key"`
	Password   string `gorm:"size:200,notnull"`
	UserInfoId int32
}

func NewUserLogin(username, password string) *UserLogin {
	return &UserLogin{
		Username: username,
		Password: password,
	}
}

// UserLoginDAO userlogin数据表操作
type UserLoginDAO struct {
}

var (
	userLoginDAO  *UserLoginDAO
	userLoginOnce sync.Once //只执行一次操作的对象
)

func NewUserLoginDAO() *UserLoginDAO {
	userLoginOnce.Do(func() {
		userLoginDAO = new(UserLoginDAO)
	})

	return userLoginDAO
}

// IsUserExistByUsername 判断用户是否存在
func (u *UserLoginDAO) IsUserExistByUsername(username string) bool {
	var login *UserLogin

	err := DB.Where("username = ?", username).Find(&login).Error
	if err != nil {
		log.Println(err)
	}

	if login.Id == 0 {
		return false
	}

	return true
}

func (u *UserLoginDAO) QueryUserLogin(username, password string)  (*UserLogin,error) {
	var login *UserLogin
	//select * from user_logins where ...
	err := DB.Where("username = ? and password = ?",username,password).Find(&login).Error
	if err != nil {
		return nil,err
	}

	return login,nil
}
