package models

import (
	"fmt"
	"testing"
)


func TestVideo(t *testing.T){
	//InitDB()
	//
	//dao := NewVideoDAO()
	//
	//video := NewVideo(2,"p","c","t")
	//
	//err := dao.AddVideo(video)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(video)


	//videos,err := dao.GetVideoListByLimitAndTime(10,time.Now())
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//for _,v := range videos{
	//	fmt.Println(v)
	//}

	//at,err := dao.GetNewVideoCreateAt()
	//if err != nil{
	//	fmt.Println(err)
	//}

	//fmt.Println(at)

	//var video *Video
	//
	////DB.First(video)
	//err := DB.Order("create_at asc").First(&video).Error
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(video)

	InitDB()

	dao := NewVideoDAO()
	//err := dao.AddVideoFavorite(1,2)
	err := dao.CancelVideoFavorite(1,2)
	video,err := dao.GetVideoByVideoId(1)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(video)
	//
	//fmt.Println("old: ",video.FavoriteCount)
	//
	//DB.Model(&video).Update("favorite_count",gorm.Expr("favorite_count - ?",1))
	//
	//time.Sleep(1*time.Second)
	//
	//fmt.Println("now: ",video.FavoriteCount)
	//DB.Model(video).Update("favorite_count",gorm.Expr("favorite_count + ?",1))
	//err != nil{
	//	fmt.Println(err)
	//}
}
