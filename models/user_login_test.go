package models

import (
	"fmt"
	"testing"
)

func TestUserLoginDAO_QueryUserLogin(t *testing.T) {
	InitDB()

	dao := NewUserLoginDAO()
	//login,err := dao.QueryUserLogin("syang","7c4a8d09ca3762af61e59520943dc26494f8941b")
	//if err != nil{
	//	fmt.Println("11111 :",err)
	//}
	//
	//fmt.Println("222   :",login)
	ok := dao.IsUserExistByUsername("syang")
	fmt.Println(ok)
}
