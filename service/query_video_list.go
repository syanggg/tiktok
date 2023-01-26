package service

import (
	"errors"
	"tiktok/cache"
	"tiktok/models"
	"tiktok/util"
	"time"
)

type FeedVideoListResponse struct {
	Videos   []*models.Video `json:"video_list"`
	NextTime int64          `json:"next_time"`		//最近投稿时间
}

func QueryFeedVideoList(userId int32,lastTime time.Time) (*FeedVideoListResponse,error){
	return NewFeedVideoFlow(userId,lastTime).Do()
}

func NewFeedVideoFlow(userId int32,lastTime time.Time) *FeedVideoListFlow{
	return &FeedVideoListFlow{
		UserId: userId,
		LastTime: lastTime,
	}
}

type FeedVideoListFlow struct {
	UserId int32
	LastTime time.Time
	*FeedVideoListResponse

	videos []*models.Video
	nextTime int64
}

func (f *FeedVideoListFlow) Do() (*FeedVideoListResponse,error){
	//参数处理
	if err := f.Check();err != nil{
		return nil,err
	}

	//数据处理
	if err := f.prepareData();err != nil{
		return nil,err
	}

	//打包response
	f.packResponse()

	return f.FeedVideoListResponse,nil
}

// Check 参数检查
func (f *FeedVideoListFlow) Check() error{
	if f.LastTime.IsZero(){
		f.LastTime = time.Now()
	}

	return nil
}

//数据处理
func (f *FeedVideoListFlow) prepareData() error{

	dao := models.NewVideoDAO()
	////用户未登录时的视频推送
	videos,err := dao.GetVideoListByLimitAndTime(util.Info.Tiktok.FeedLimit,time.Now())
	if err != nil{
		return err
	}


	//填充作者信息
	if err := fillVideoAuthor(videos);err != nil{
		return err
	}


	//todo 用户登录时的视频推送

	//用户已登录更新获取视频的关注状态和点赞状态
	if f.UserId != 0{
		err := fillVideoFavoriteAndFollow(f.UserId,videos)
		if err != nil{
			return err
		}
	}

	//对videos处理后赋值
	f.videos = videos

	//返回最新投稿时间
	newTime,err := dao.GetNewVideoCreateAt()
	if err != nil{
		return err
	}

	f.nextTime = newTime * 1e3


	return nil
}

//打包response
func (f *FeedVideoListFlow) packResponse(){
	f.FeedVideoListResponse = &FeedVideoListResponse{
		Videos: f.videos,
		NextTime: f.nextTime,
	}
}


//填充视频点赞和关注信息
func fillVideoFavoriteAndFollow(userId int32,videos []*models.Video) error{
	redis := cache.NewRedisOperation()

	for _,video := range videos{
		//填充点赞状态
		isFavorite,err :=  redis.GetFavoriteState(userId,video.Id)
		if err != nil{
			return err
		}
		video.IsFavorite = isFavorite
		//填充关注状态
		isFollow,err := redis.GetFollowState(userId,video.UserInfoId)
		if err != nil{
			return err
		}
		if video.Author == nil{
			return errors.New("video author info not filled")
		}

		video.Author.IsFollow = isFollow
	}


	return nil
}



//FillVideoAuthor 填充视频作者信息
func fillVideoAuthor(videos []*models.Video) error{
	dao := models.NewUserInfoDAO()

	for _,video := range videos{
		info,err := dao.GetUserInfoById(video.UserInfoId)
		if err != nil{
			return err
		}

		video.Author = info
	}

	return nil
}
