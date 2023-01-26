package service

import (
	"fmt"
	"testing"
	"tiktok/models"
	"time"
)

func TestQueryFeedVideoList(t *testing.T) {
	models.InitDB()
	response,err := QueryFeedVideoList(0,time.Now())
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(response.Videos)
	fmt.Println(response.NextTime)

}