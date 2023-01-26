package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/models"
	"tiktok/service"
)

type LoginResponse struct {
	models.CommonResponse
	*service.LoginInfoResponse
}

func LoginHandle(c *gin.Context){
	username := c.Query("username")
	any,_ := c.Get("password")
	password ,ok:= any.(string)
	if !ok {
		c.JSON(http.StatusOK,&LoginResponse{
			CommonResponse:models.CommonResponse{
				StatusCode: 1,
				StatusMsg: "failed to decode password",
			},
		})
	}

	response,err := service.QueryLogin(username,password)
	if err != nil{
		c.JSON(http.StatusOK,&LoginResponse{
			CommonResponse:models.CommonResponse{
				StatusCode: 1,
				StatusMsg: err.Error(),
			},
		})
	}

	c.JSON(http.StatusOK,&LoginResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0},
		LoginInfoResponse: response,
	})

}
