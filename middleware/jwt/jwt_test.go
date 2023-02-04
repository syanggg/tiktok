package middleware

import (
	"fmt"
	"testing"
	"tiktok/middleware/jwt"
	"tiktok/models"
)

func TestJwtMiddleware(t *testing.T) {
	login := &models.UserLogin{UserInfoId: 10}

	token,err := jwt.ReleaseToken(login)
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(token)

	claims,err := jwt.ParseToken(token)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(claims)
}
