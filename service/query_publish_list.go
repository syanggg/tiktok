package service

import (
	"errors"
	"tiktok/models"
)

type PublishVideoListResponse struct {
	Videos []*models.Video `json:"video_list,omitempty"`
}

func QueryPublishVideoList(userId int32) (*PublishVideoListResponse,error){
	return NewPublishVideoListFlow(userId).Do()
}

func NewPublishVideoListFlow(userId int32) *PublishVideoListFlow{
	return &PublishVideoListFlow{
		UserId: userId,
	}
}

type PublishVideoListFlow struct {
	UserId int32

	*PublishVideoListResponse
	videos []*models.Video
}

func (p *PublishVideoListFlow) Do() (*PublishVideoListResponse,error){
	if err := p.Check();err != nil{
		return nil,err
	}

	err := p.prepareData()
	if err != nil{
		return nil,err
	}

	p.packResponse()

	return p.PublishVideoListResponse,nil
}

// Check 参数检查
func (p *PublishVideoListFlow) Check() error{
	if p.UserId == 0{
		return errors.New("user not login")
	}

	return nil
}

//数据处理
func (p *PublishVideoListFlow) prepareData() error{
	dao := models.NewVideoDAO()
	//通过用户id获取发布视频列表
	list,err := dao.GetVideoListByUserId(p.UserId)
	if err != nil{
		return err
	}

	p.videos = list

	return nil
}

//打包response
func (p *PublishVideoListFlow) packResponse(){
	p.PublishVideoListResponse = &PublishVideoListResponse{
		Videos: p.videos,
	}
}