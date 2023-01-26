package service

import (
	"fmt"
	"testing"
	"tiktok/cache"
	"tiktok/models"
)

func TestFavoriteFlow(t *testing.T) {
	cache.InitRedis()
	models.InitDB()

	//flow := NewFavoriteFlow(1,1,2)
	//err := flow.Do()
	//if err != nil{
	//	fmt.Println(err)
	//}

	op := cache.NewRedisOperation()
	list,err := op.GetFavoriteVideoIdList(2)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(list)
}
