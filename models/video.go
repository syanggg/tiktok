package models

import (
	"gorm.io/gorm"
	"log"
	"sync"
	"time"
)

func NewVideo(userId int32,playUrl string,coverUrl string,title string) *Video{
	return &Video{
		UserInfoId: userId,
		PlayUrl: playUrl,
		CoverUrl: coverUrl,
		Title: title,
		CreateAt: time.Now(),
	}
}

type Video struct {
	Id            int32       `json:"id"`
	UserInfoId    int32       `json:"-"`	//作者id 一对多关系(查询一般在多的表操作)
	Author        *UserInfo   `json:"author,omitempty" gorm:"-"`	//作者信息
	PlayUrl       string      `json:"play_url,omitempty"`		//视频url
	CoverUrl      string      `json:"cover_url,omitempty"`		//封面url
	FavoriteCount int32       `json:"favorite_count,omitempty"`	//喜欢数量
	CommentCount  int32       `json:"comment_count,omitempty"`	//评论数量
	IsFavorite    bool        `json:"is_favorite,omitempty"`	//是否喜欢
	Title         string      `json:"title,omitempty"`			//标题
	Users         []*UserInfo `json:"-" gorm:"many2many:user_favor_videos;"`	//多对多关系 视频与喜欢的人
	Comment       []*Comment    `json:"-"`	//评论 一对多关系（不会被gorm建表）
	CreateAt      time.Time   `json:"-"`
}


type VideoDAO struct {
}

var (
	videoDAO *VideoDAO
	videoOnce sync.Once
)

func NewVideoDAO() *VideoDAO{
	videoOnce.Do(func() {
		videoDAO = new(VideoDAO)
	})

	return videoDAO
}

// GetVideoListByUserId 通过用户id获取用户发布的视频列表
func (v *VideoDAO) GetVideoListByUserId(id int32) ([]*Video,error){
	var list []*Video
	//select * from videos where user_info_id = 'id'
	err := DB.Where("user_info_id = ?",id).Find(&list).Error
	if err != nil{
		return nil,err
	}

	return list,nil
}

// GetVideoCountByUserId 通过用户id获取用户发布的视频数量
func (v *VideoDAO) GetVideoCountByUserId(id int32) (int64,error){
	var count int64
	err := DB.Where("user_info_id = ?",id).Count(&count).Error
	if err != nil{
		return 0,err
	}

	return count,err
}

// GetVideoListByLimitAndTime 通过时间戳查询时间戳之前的limit个视频列表
func (v *VideoDAO) GetVideoListByLimitAndTime(limit int,time time.Time) ([]*Video,error){
	var list []*Video
	//select * from videos order by create_at asc limit 'limit'
	err := DB.Where("create_at < ?",time).Order("create_at asc").Limit(limit).Find(&list).Error
	if err != nil{
		return nil,err
	}

	return list,nil
}

// GetVideoByVideoId 通过视频Id获取视频
func (v *VideoDAO) GetVideoByVideoId(id int32) (*Video,error) {
	var video *Video

	err := DB.Where("id = ?",id).Find(&video).Error
	if err != nil{
		return nil,err
	}

	return video,nil
}

// GetNewVideoCreateAt 获取最近投稿时间(s时间戳)
func (v *VideoDAO) GetNewVideoCreateAt() (int64,error){
	var video *Video
	err := DB.Order("create_at asc").First(&video).Error
	if err != nil{
		return 0,err
	}

	newTime := video.CreateAt.Unix()

	return newTime,nil
}



// AddVideo 添加视频
func (v *VideoDAO) AddVideo(video *Video) error{
	if video == nil{
		return ErrNilPtr
	}

	err := DB.Create(video).Error
	if err != nil{
		return err
	}

	return nil
}

// AddVideoFavorite 添加一个赞
func (v *VideoDAO) AddVideoFavorite(videoId int32, userId int32) error{
	video,err := v.GetVideoByVideoId(videoId)
	if err != nil{
		return err
	}
	info,err := NewUserInfoDAO().GetUserInfoById(userId)
	if err != nil{
		return err
	}


	if err := DB.Model(video).Update("favorite_count",gorm.Expr("favorite_count + ?",1)).Error;err != nil{
		return err
	}

	//添加关联
	if err := DB.Model(video).Association("Users").Append(info);err != nil{
		return err
	}


	return nil
}

// CancelVideoFavorite 取消一个赞
func (v *VideoDAO) CancelVideoFavorite(videoId int32, userId int32) error{
	video,err := v.GetVideoByVideoId(videoId)
	if err != nil{
		return err
	}
	info,err := NewUserInfoDAO().GetUserInfoById(userId)
	if err != nil{
		return err
	}

	//减少赞的数量
	//update videos set favorite_count = favorite_count + 1
	if err := DB.Model(video).Where("favorite_count > ?",0).Update("favorite_count",
		gorm.Expr("favorite_count - ?",1)).Error;err != nil{
		return nil
	}

	//删除关联
	if err := DB.Model(video).Association("Users").Delete(info);err != nil{
		return err
	}

	return nil
}

// AddVideoComment 增加一个评论数量
func (v *VideoDAO) AddVideoComment(videoId int32) error{
	video,err := v.GetVideoByVideoId(videoId)
	if err != nil{
		return err
	}

	if err := DB.Model(video).Update("comment_count",gorm.Expr("comment_count + ?",1)).Error;err != nil{
		return err
	}

	return nil
}

// CancelVideoComment 减少一个评论数量
func (v *VideoDAO) CancelVideoComment(videoId int32) error{
	video,err := v.GetVideoByVideoId(videoId)
	if err != nil{
		return err
	}

	if err := DB.Model(video).Update("comment_count",gorm.Expr("comment_count - ?",1)).Error;err != nil{
		return err
	}

	return nil
}

// IsVideoExitsByVideoId 通过id判断视频是否存在
func (v *VideoDAO) IsVideoExitsByVideoId(id int32) bool{
	var video Video

	err := DB.Where("id = ?",id).Find(&video).Error
	if err != nil{
		log.Println(err)
	}
	if video.Id == 0 {
		return false
	}

	return true

}