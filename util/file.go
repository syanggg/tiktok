package util

import (
	"fmt"
	"log"
	"tiktok/models"
)

// FileName 拼接filename
func FileName(userId int32,suffix string) string{
	dao := models.NewVideoDAO()

	count,err := dao.GetVideoCountByUserId(userId)
	if err != nil{
		log.Println(err.Error())
	}

	name := fmt.Sprintf("%d-%d%s",userId,count,suffix)

	return name
}


func GetFileUrl(filename string) string{
	url := fmt.Sprintf("http://%s:%d/static/%s", Info.Server.Host, Info.Server.Port,filename)

	return url
}