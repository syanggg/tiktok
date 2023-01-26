package service

import (
	"fmt"
	"testing"
	"tiktok/models"
)

func TestQueryLogin(t *testing.T) {
	models.InitDB()

	response,err := QueryLogin("user","123456")
	if err != nil{
		fmt.Println(err)
	}

	fmt.Println(response)
}
