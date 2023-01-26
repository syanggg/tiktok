package middleware

import (
	"fmt"
	"testing"
	"tiktok/models"
)

func TestJwtMiddleware(t *testing.T) {
	login := &models.UserLogin{UserInfoId: 10}

	token,err := ReleaseToken(login)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(token)

	claims,err := ParseToken(token)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(claims)
}
