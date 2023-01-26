package service

import (
	"errors"
	"tiktok/models"
)

const (
	CREATE = 1
	DELETE = 2
)

type CommentInfoResponse struct {
	Comment *models.Comment `json:"comment"`
}

func PostComment(userId int32, videoId int32, actionType int32, commentId int32, content string) (*CommentInfoResponse, error) {
	return NewCommentFlow(userId, videoId, actionType, commentId, content).Do()
}

func NewCommentFlow(userId int32, videoId int32, actionType int32, commentId int32, content string) *CommentInfoFlow {
	return &CommentInfoFlow{
		UserId:     userId,
		VideoId:    videoId,
		ActionType: actionType,
		CommentId:  commentId,
		Content:    content,
	}
}

type CommentInfoFlow struct {
	UserId     int32
	VideoId    int32
	ActionType int32
	CommentId  int32
	Content    string

	*CommentInfoResponse
	comment *models.Comment
}

func (c *CommentInfoFlow) Do() (*CommentInfoResponse, error) {
	if err := c.Check(); err != nil {
		return nil, err
	}

	if err := c.prepareData(); err != nil {
		return nil, err
	}

	c.packResponse()

	return c.CommentInfoResponse, nil
}

// Check 参数检查
func (c *CommentInfoFlow) Check() error {
	if !models.NewUserInfoDAO().IsUserExitsById(c.UserId) {
		return errors.New("user not found")
	}
	if !models.NewVideoDAO().IsVideoExitsByVideoId(c.VideoId) {
		return errors.New("video not found")
	}
	if c.ActionType != CREATE && c.ActionType != DELETE {
		return errors.New("actionType err")
	}

	return nil
}

//数据处理
func (c *CommentInfoFlow) prepareData() error {
	switch c.ActionType {
	case CREATE:
		if err := c.CreateComment(); err != nil {
			return err
		}
	case DELETE:
		if err := c.DeleteComment(); err != nil {
			return err
		}
	default:
		return errors.New("actionType err")
	}




	return nil
}

//打包response
func (c *CommentInfoFlow) packResponse() {
	c.CommentInfoResponse = &CommentInfoResponse{
		Comment: c.comment,
	}
}

// CreateComment 添加评论
func (c *CommentInfoFlow) CreateComment() error {
	dao := models.NewCommentDAO()
	comment := models.NewComment(c.UserId, c.VideoId, c.Content)

	err := dao.AddComment(comment)
	if err != nil {
		return err
	}

	if err := models.NewVideoDAO().AddVideoComment(c.VideoId);err != nil{
		return err
	}

	//填充信息
	if err := fillCommentInfo(comment);err != nil{
		return err
	}

	c.comment = comment

	return nil
}

// DeleteComment 删除评论
func (c *CommentInfoFlow) DeleteComment() error {
	dao := models.NewCommentDAO()
	comment, err := dao.GetCommentByCommentId(c.CommentId)
	if err != nil {
		return err
	}

	if err := dao.DeleteComment(comment); err != nil {
		return err
	}

	if err := models.NewVideoDAO().CancelVideoComment(c.VideoId);err != nil{
		return err
	}

	return nil
}

//填充时间信息
func fillCommentInfo(comment *models.Comment) error{
	if comment == nil{
		return errors.New("comment nil")
	}

	comment.CreateDate = comment.CreateAt.Format("1-2")

	return nil
}