package service

import (
	"errors"
	"tiktok/models"
)

type CommentInfoListResponse struct {
	Comments []*models.Comment `json:"comment_list"`
}

func QueryCommentList(userId int32,videoId int32) (*CommentInfoListResponse,error){
	return NewCommentInfoListFlow(userId,videoId).Do()
}

func NewCommentInfoListFlow(userId int32,videoId int32) *CommentInfoListFlow{
	return &CommentInfoListFlow{
		UserId: userId,
		VideoId: videoId,
	}
}


type CommentInfoListFlow struct {
	UserId int32
	VideoId int32

	*CommentInfoListResponse
	comments []*models.Comment
}

func (c *CommentInfoListFlow) Do() (*CommentInfoListResponse,error){
	//参数检查
	if err := c.Check();err != nil{
		return nil,err
	}

	//数据处理
	if err := c.prepareData();err != nil{
		return nil,err
	}

	//打包response
	c.packResponse()

	return c.CommentInfoListResponse,nil
}

// Check 参数检查
func (c *CommentInfoListFlow) Check() error{
	if !models.NewUserInfoDAO().IsUserExitsById(c.UserId){
		return errors.New("user not found")
	}

	if !models.NewVideoDAO().IsVideoExitsByVideoId(c.VideoId){
		return errors.New("video not found")
	}

	return nil
}

//数据处理
func (c *CommentInfoListFlow) prepareData() error{
	dao := models.NewCommentDAO()

	comments,err := dao.GetCommentListByVideoId(c.VideoId)
	if err != nil {
		return err
	}

	//填充信息
	if err := fillCommentsInfo(comments);err != nil{
		return err
	}

	c.comments = comments

	return nil
}

//打包response
func (c *CommentInfoListFlow) packResponse(){
	c.CommentInfoListResponse = &CommentInfoListResponse{
		Comments: c.comments,
	}
}

//填充comment信息
func fillCommentsInfo(comments []*models.Comment) error{
	if comments == nil{
		return errors.New("comments nil")
	}

	for _,comment := range comments{
		info,err := models.NewUserInfoDAO().GetUserInfoById(comment.UserInfoId)
		if err != nil{
			return err
		}
		comment.User = info

		comment.CreateDate = comment.CreateAt.Format("1-2")

	}

	return nil
}