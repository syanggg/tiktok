package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/models"
	"tiktok/service"
)

type FavoriteListResponse struct {
	models.CommonResponse
	*service.FavoriteVideoListResponse
}

func FavoriteListHandle(c *gin.Context){
	any,_ := c.Get("user_id")
	userId,ok := any.(int32)
	if !ok{
		favoriteListResponseErr(c,"failed to decode userId")
		return
	}

	//service层处理
	response, err := service.QueryFavoriteVideoList(userId)
	if err != nil {
		favoriteListResponseErr(c,err.Error())
		return
	}

	favoriteListResponseOk(c,response)
}


func favoriteListResponseErr(c *gin.Context,msg string){
	c.JSON(http.StatusOK,FavoriteListResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 1,
			StatusMsg: msg,
		},
	})
}

func favoriteListResponseOk(c *gin.Context,response *service.FavoriteVideoListResponse){
	c.JSON(http.StatusOK,FavoriteListResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 0,
		},
		FavoriteVideoListResponse: response,
	})
}