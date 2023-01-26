package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/models"
	"tiktok/service"
)

type FollowersResponse struct {
	models.CommonResponse
	*service.FollowerListResponse
}

func FollowerListHandle(c *gin.Context){
	any,_ := c.Get("user_id")
	userId,ok := any.(int32)
	if !ok {
		followersResponseErr(c,"failed to decode userId")
		return
	}

	response,err := service.QueryFollowerList(userId)
	if err != nil{
		followersResponseErr(c,err.Error())
		return
	}

	followersResponseOk(c,response)
}


func followersResponseOk(c *gin.Context,response *service.FollowerListResponse){
	c.JSON(http.StatusOK,FollowersResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 0,
		},
		FollowerListResponse:response,
	})
}

func followersResponseErr(c *gin.Context,msg string){
	c.JSON(http.StatusOK, FollowersResponse{
			CommonResponse:models.CommonResponse{
				StatusCode: 0,
				StatusMsg: msg,
			},
	})
}