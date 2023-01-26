package models

import (
	"sync"
	"time"
)

func NewComment(userId int32, videoId int32, content string) *Comment {
	return &Comment{
		UserInfoId: userId,
		VideoId:    videoId,
		Content:    content,
		CreateAt:   time.Now(),
	}
}

type Comment struct {
	Id         int32     `json:"id"`
	UserInfoId int32     `json:"-"`
	VideoId    int32     `json:"-"`
	User       *UserInfo  `json:"user" gorm:"-"`
	Content    string    `json:"content"`
	CreateAt   time.Time `json:"-"`
	CreateDate string    `json:"create_date" gorm:"-"` //前端要求的日期格式
}

type CommentDAO struct {
}

var (
	commentDAO  *CommentDAO
	commentOnce sync.Once
)

func NewCommentDAO() *CommentDAO {
	commentOnce.Do(func() {
		commentDAO = new(CommentDAO)
	})

	return commentDAO
}

// GetCommentListByVideoId 通过视频Id返回评论列表
func (c *CommentDAO) GetCommentListByVideoId(videoId int32) ([]*Comment, error) {
	var list []*Comment

	err := DB.Where("video_id = ?", videoId).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// GetCommentListByUserId 通过用户Id返回评论列表
func (c *CommentDAO) GetCommentListByUserId(userId int32) ([]*Comment, error) {
	var list []*Comment

	err := DB.Where("user_info_id = ?", userId).Find(&list).Error
	if err != nil {
		return nil, err
	}

	return list, nil
}

// AddComment 添加评论
func (c *CommentDAO) AddComment(comment *Comment) error {
	if comment == nil {
		return ErrNilPtr
	}

	err := DB.Create(comment).Error
	if err != nil {
		return err
	}

	return nil
}

// DeleteComment 删除评论
func (c *CommentDAO) DeleteComment(comment *Comment) error {
	if comment == nil {
		return ErrNilPtr
	}

	err := DB.Delete(comment).Error
	if err != nil {
		return err
	}
	return nil
}

// GetCommentByCommentId 通过评论id获取评论
func (c *CommentDAO) GetCommentByCommentId(commentId int32) (*Comment, error) {
	var comment *Comment

	err := DB.Where("id = ?", commentId).Find(&comment).Error
	if err != nil {
		return nil, err
	}

	return comment, nil
}
