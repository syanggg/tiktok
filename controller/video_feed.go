package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/models"
	"tiktok/service"
	"time"
)

type FeedResponse struct {
	models.CommonResponse
	*service.FeedVideoListResponse
}

func FeedVideoListHandle(c *gin.Context){
	rawTime := c.Query("latest_time")

	timeInt,err := strconv.ParseInt(rawTime,10,64)
	if err != nil{
		timeInt = time.Now().UnixMilli()
	}

	lastTime := time.Unix(0,timeInt*1e6) //前端传来的是ms


	//todo 通过token判断登录状态
	//any,ok := c.GetQuery("token")  //GetQuery()相比Query()多了一个判断是否获取成功返回

	//用户未登录时 feedMiddleware中间件不提供user_id
	any,ok := c.Get("user_id")
	//未登录
	if !ok {
		response,err := service.QueryFeedVideoList(0,lastTime)
		if err != nil{
			responseErr(c,err.Error())
			return
		}

		responseOk(c,response)

		return
	}


	//todo 已登录时存在bug
	//已登录
	id,ok := any.(int32)
	if !ok{
		responseErr(c,"failed to decode userid")
		return
	}
	response,err := service.QueryFeedVideoList(id,lastTime)
	if err != nil{
		responseErr(c,err.Error())
		return
	}

	responseOk(c,response)

}

func responseOk(c *gin.Context,response *service.FeedVideoListResponse){
	c.JSON(http.StatusOK,FeedResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 0,
		},
		FeedVideoListResponse:response,
	})
}

func responseErr(c *gin.Context,msg string){
	c.JSON(http.StatusOK,FeedResponse{
		CommonResponse:models.CommonResponse{
			StatusCode: 1,
			StatusMsg: msg,
		},
	})
}