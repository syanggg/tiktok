package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"tiktok/models"
	"tiktok/service"
	"tiktok/util"
	"time"
)

type PublishResponse struct {
	models.CommonResponse
}

var videoSuffixs = map[string]struct{}{
	".mp4":  {},
	".avi":  {},
	".wmv":  {},
	".flv":  {},
	".mpeg": {},
	".mov":  {},
}

func PublishHandle(c *gin.Context){
	any,_ := c.Get("user_id")

	id,ok := any.(int32)
	if !ok{
		publishResponseErr(c,"failed to decode userid")
		return
	}

	title := c.PostForm("title")


	form,err := c.MultipartForm()  //解析表单
	if err != nil{
		publishResponseErr(c,err.Error())
		return
	}
	files := form.File["data"] //返回指定表单键的文件map集合
	//c.FormFile() //返回指定表单键对应的第一个文件

	for _,file := range files{
		suffix := filepath.Ext(file.Filename)
		if _,ok := videoSuffixs[suffix];!ok{
			publishResponseErr(c,"failed to file publish,suffix unsupported")
			return
		}
		//根据userid拼接文件名称
		fileName := util.FileName(id,suffix)
		//拼接视频保存位置
		savePath := filepath.Join("./static",fileName)

		//保存
		err := c.SaveUploadedFile(file,savePath)
		if err != nil{
			publishResponseErr(c,err.Error())
			return
		}

		//封面
		coverName := util.FileName(id,".jpg")

		bind := util.NewBind(util.Info.Path.FfmpegPath)
		input := filepath.Join(util.Info.Path.StaticPath,fileName)
		output := filepath.Join(util.Info.Path.StaticPath,coverName)
		if err := bind.Thumbnail(input,output,1*time.Second,false);err != nil{
			publishResponseErr(c,err.Error())
			return
		}


		//数据库操作
		err = service.PublishVideo(id,fileName,coverName,title)
		if err != nil{
			publishResponseErr(c,err.Error())
			return
		}

		publishResponseOk(c,"publish success")

	}

}


func publishResponseErr(c *gin.Context,msg string){
	c.JSON(http.StatusOK,PublishResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 1,
			StatusMsg: msg,
		},
	})
}

func publishResponseOk(c *gin.Context,msg string){
	c.JSON(http.StatusOK,PublishResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 0,
			StatusMsg: msg,
		},
	})
}