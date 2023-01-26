package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/models"
	"tiktok/service"
)

type RegisterResponse struct {
	models.CommonResponse
	*service.LoginInfoResponse
}

func RegisterHandle(c *gin.Context){
	username := c.Query("username")
	any,_ := c.Get("password")
	password,ok := any.(string)
	if !ok {
		c.JSON(http.StatusOK,RegisterResponse{
			CommonResponse:models.CommonResponse{
				StatusCode: 1,		//code 1表示错误
				StatusMsg: "密码解析失败",
			},
		})
		return
	}

	response,err := service.PostRegister(username,password)
	if err != nil{
		c.JSON(http.StatusOK,RegisterResponse{
			CommonResponse:models.CommonResponse{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
		})
		return
	}

	c.JSON(http.StatusOK,RegisterResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0}, //code 0 表示成功
		LoginInfoResponse: response,
	})

}

