package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tiktok/models"
	"tiktok/service"
)

type PublishListResponse struct {
	models.CommonResponse
	*service.PublishVideoListResponse
}

func PublishListHandle(c *gin.Context){
	any,_ := c.Get("user_id")
	id,ok := any.(int32)
	if !ok{
		publishResponseErr(c,"failed to decode userid")
	}

	//service层对数据库操作
	response,err := service.QueryPublishVideoList(id)
	if err != nil{
		publishListResponseErr(c,err.Error())
	}

	publishListResponseOk(c,response)

}


func publishListResponseOk(c *gin.Context,response *service.PublishVideoListResponse){
	c.JSON(http.StatusOK,PublishListResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 0,
		},
		PublishVideoListResponse:response,
	})
}

func publishListResponseErr(c *gin.Context,msg string){
	c.JSON(http.StatusOK,PublishListResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 1,
			StatusMsg: msg,
		},
	})
}