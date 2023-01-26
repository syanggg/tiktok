package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/models"
	"tiktok/service"
)

type FollowResponse struct {
	models.CommonResponse
}

func FollowHandle(c *gin.Context){
	any,_ := c.Get("user_id")
	userId,ok := any.(int32)
	if !ok{
		followResponseErr(c,"failed to decode userId")
		return
	}

	rawFollowId := c.Query("to_user_id")
	followId,err := strconv.ParseInt(rawFollowId,10,64)
	if err != nil{
		followResponseErr(c,err.Error())
		return
	}

	rawActionType := c.Query("action_type")
	actionType,err := strconv.ParseInt(rawActionType,10,64)
	if err != nil{
		followResponseErr(c,err.Error())
		return
	}

	if err := service.PostFollow(userId,int32(followId),int32(actionType));err != nil{
		followResponseErr(c,err.Error())
		return
	}

	followResponseOk(c,"follow success")

}

func followResponseOk(c *gin.Context,msg string){
	c.JSON(http.StatusOK,FollowResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 0,
			StatusMsg: msg,
		},
	})
}

func followResponseErr(c *gin.Context,msg string){
	c.JSON(http.StatusOK,FollowResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 1,
			StatusMsg: msg,
		},
	})
}
