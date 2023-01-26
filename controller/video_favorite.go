package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/models"
	"tiktok/service"
)

type FavoriteResponse struct {
	models.CommonResponse
}

func FavoriteHandle(c *gin.Context){
	any,_ := c.Get("user_id")
	userId,ok := any.(int32)
	if !ok{
		favoriteResponseErr(c,"failed to decode userid")
		return
	}
	rawVideoId := c.Query("video_id")
	videoId,err := strconv.ParseInt(rawVideoId,10,64)
	if err != nil{
		responseErr(c,err.Error())
		return
	}
	rawActionType := c.Query("action_type")
	actionType,err := strconv.ParseInt(rawActionType,10,64)
	if err != nil{
		responseErr(c,err.Error())
		return
	}

	//数据库操作
	if err := service.PostFavorite(userId, int32(videoId), int32(actionType));err != nil {
		favoriteResponseErr(c,err.Error())
		return
	}

	favoriteResponseOk(c)
}


func favoriteResponseOk(c *gin.Context){
	c.JSON(http.StatusOK,FavoriteResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 0,
		},
	})
}

func favoriteResponseErr(c *gin.Context,msg string){
	c.JSON(http.StatusOK,FavoriteResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 1,
			StatusMsg: msg,
		},
	})
}