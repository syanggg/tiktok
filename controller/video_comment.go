package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/models"
	"tiktok/service"
)

type CommentResponse struct {
	models.CommonResponse
	*service.CommentInfoResponse
}

func CommentHandle(c *gin.Context) {
	any, _ := c.Get("user_id")
	userId, ok := any.(int32)
	if !ok {
		commentResponseErr(c, "failed to decode userId")
		return
	}

	rawVideoId := c.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		commentResponseErr(c, err.Error())
		return
	}

	rawActionType := c.Query("action_type")
	actionType, err := strconv.ParseInt(rawActionType, 10, 64)
	if err != nil {
		commentResponseErr(c, err.Error())
		return
	}

	var content string
	var commentId int64

	switch actionType {
	case service.CREATE:
		content = c.Query("comment_text")
	case service.DELETE:
		rawCommentId := c.Query("comment_id")
		commentId, err = strconv.ParseInt(rawCommentId, 10, 64)
		if err != nil {
			commentResponseErr(c, err.Error())
			return
		}
	default:
		commentResponseErr(c, "actionType err")
	}

	response, err := service.PostComment(userId, int32(videoId), int32(actionType), int32(commentId), content)
	if err != nil {
		commentResponseErr(c, err.Error())
		return
	}

	commentResponseOk(c, response)
}

func commentResponseErr(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, CommentResponse{
		CommonResponse: models.CommonResponse{
			StatusCode: 1,
			StatusMsg:  msg,
		},
	})
}

func commentResponseOk(c *gin.Context, response *service.CommentInfoResponse) {
	c.JSON(http.StatusOK, CommentResponse{
		CommonResponse: models.CommonResponse{
			StatusCode: 0,
		},
		CommentInfoResponse: response,
	})
}
