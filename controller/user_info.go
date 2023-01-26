package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/models"
	"tiktok/service"
)

type UserResponse struct {
	models.CommonResponse
	*service.UserInfoResponse
}

func UserInfoHandle(c *gin.Context){
	any,_ := c.Get("user_id")

	id,ok := any.(int32)
	if !ok {
		c.JSON(http.StatusOK,UserResponse{
			CommonResponse:models.CommonResponse{
				StatusCode: 1,
				StatusMsg: "failed to parse userid",
			},
		})
	}

	response,err := service.QueryUserInfo(id)
	if err != nil{
		c.JSON(http.StatusOK,UserResponse{
			CommonResponse:models.CommonResponse{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
		})
	}

	c.JSON(http.StatusOK,UserResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 0,
		},
		UserInfoResponse: response,
	})

}