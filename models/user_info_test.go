package models

import (
	"fmt"
	"testing"
)

func TestUserInfoDAO_AddUserFollow(t *testing.T) {

	InitDB()

	userInfoDAO := NewUserInfoDAO()

	xiaop := NewUserInfo("xiaop",&UserLogin{})
	xiaot := NewUserInfo("xiaot",&UserLogin{})


	err := userInfoDAO.AddUserInfo(xiaop)
	if err != nil{
		fmt.Println(err)
	}
	err = userInfoDAO.AddUserInfo(xiaot)
	if err != nil{
		fmt.Println(err)
	}



	err = userInfoDAO.AddUserFollow(xiaop.Id,xiaot.Id)
	if err != nil{
		fmt.Println(err)
	}


}

func TestUserInfoDAO_CancelUserFollow(t *testing.T) {
	InitDB()

	userInfoDAO := NewUserInfoDAO()

	//syang := &UserInfo{}
	//xiaok := &UserInfo{}
	//xiaof := &UserInfo{}

	xiaop,_ := userInfoDAO.GetUserInfoById(7)
	xiaot,_ := userInfoDAO.GetUserInfoById(7)

	userInfoDAO.CancelUserFollow(xiaop.Id,xiaot.Id)


}

func TestUserInfoDAO_GetFollowListById(t *testing.T) {
	InitDB()

	userInfoDAO := NewUserInfoDAO()

	syang,_ := userInfoDAO.GetUserInfoById(1)
	//xiaok,_ := userInfoDAO.GetUserInfoById(2)
	//xiaof,_ := userInfoDAO.GetUserInfoById(3)

	//userInfoDAO.AddUserFollow(xiaok.Id,syang.Id)
	//userInfoDAO.AddUserFollow(xiaof.Id,syang.Id)
	//userInfoDAO.CancelUserFollow(2,1)

	list,err := userInfoDAO.GetFollowListById(syang.Id)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println("len: ",len(list),"follow:")
	for _,info := range list{
		fmt.Println(info.Id)
	}
	//info,err := userInfoDAO.GetFollowListById(syang.Id)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Println("id:",info.Id,"name:",info.Name,info.FollowCount,info.FollowerCount,len(info.Follows))
}

func TestUserInfoDAO_GetFollowerListById(t *testing.T) {
	InitDB()

	userInfoDAO := NewUserInfoDAO()

	//syang,err := userInfoDAO.GetUserInfoById(1)
	//if err != nil{
	//	fmt.Println(err)
	//}

	//userInfoDAO.AddUserFollow(2,1)

	list,err := userInfoDAO.GetFollowerListById(1)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println("len: ",len(list),"follower:")
	for _,info := range list{
		fmt.Println(info.Id)
	}


}

func TestUserInfoDAO_IsUserExitsById(t *testing.T) {
	InitDB()

	dao := NewUserInfoDAO()

	dao.IsUserExitsById(1)
}