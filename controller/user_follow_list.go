package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/models"
	"tiktok/service"
)

type FollowsResponse struct {
	models.CommonResponse
	*service.FollowListResponse
}

func FollowListHandle(c *gin.Context){
	any,_ := c.Get("user_id")
	userId,ok := any.(int32)
	if !ok {
		followsResponseErr(c,"failed to decode userId")
		return
	}

	response,err := service.QueryFollowList(userId)
	if err != nil{
		followsResponseErr(c,err.Error())
		return
	}

	followsResponseOk(c,response)
}


func followsResponseOk(c *gin.Context,response *service.FollowListResponse){
	c.JSON(http.StatusOK,FollowsResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 0,
		},
		FollowListResponse: response,
	})
}

func followsResponseErr(c *gin.Context,msg string){
	c.JSON(http.StatusOK,FollowsResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 1,
			StatusMsg: msg,
		},
	})
}