package service

import (
	"errors"
	"tiktok/models"
	"tiktok/util"
)

func PublishVideo(userId int32,playName,coverName,title string) error {
	playUrl := util.GetFileUrl(playName)
	coverUrl := util.GetFileUrl(coverName)

	return	NewPublishFlow(userId,playUrl,coverUrl,title).Do()
}

func NewPublishFlow(userId int32,playUrl,coverUrl,title string) *PublishFlow{
	return &PublishFlow{
		UserId: userId,
		PlayUrl: playUrl,
		CoverUrl: coverUrl,
		Title: title,
	}
}

type PublishFlow struct {
	UserId int32
	PlayUrl string
	CoverUrl string
	Title string
}


func (p *PublishFlow) Do() error{
	if err := p.Check();err != nil{
		return err
	}

	err := p.prepareData()
	if err != nil{
		return err
	}

	return nil
}

// Check 参数检查
func (p *PublishFlow) Check() error{
	if p.UserId == 0{
		return errors.New("user not login")
	}

	return nil
}

//数据处理
func (p *PublishFlow) prepareData() error {
	dao := models.NewVideoDAO()

	video := models.NewVideo(p.UserId,p.PlayUrl,p.CoverUrl,p.Title)

	err := dao.AddVideo(video)
	if err != nil {
		return err
	}

	return nil
}
