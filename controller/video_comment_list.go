package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/models"
	"tiktok/service"
)

type CommentListResponse struct{
	models.CommonResponse
	*service.CommentInfoListResponse
}


func CommentListHandle(c *gin.Context){
	any,_ := c.Get("user_id")
	userId,ok := any.(int32)
	if !ok{
		commentListResponseErr(c,"failed to decode userId")
		return
	}

	rawVideoId := c.Query("video_id")
	videoId,err := strconv.ParseInt(rawVideoId,10,64)
	if err != nil{
		commentListResponseErr(c,err.Error())
		return
	}

	response,err := service.QueryCommentList(userId,int32(videoId))
	if err != nil{
		commentListResponseErr(c,err.Error())
		return
	}

	commentListResponseOk(c,response)

}

func commentListResponseOk(c *gin.Context,response *service.CommentInfoListResponse){
	c.JSON(http.StatusOK,CommentListResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 0,
		},
		CommentInfoListResponse: response,
	})
}

func commentListResponseErr(c *gin.Context,msg string){
	c.JSON(http.StatusOK,CommentListResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 1,
			StatusMsg: msg,
		},
	})
}