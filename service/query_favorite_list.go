package service

import (
	"errors"
	"strconv"
	"tiktok/cache"
	"tiktok/models"
)

type FavoriteVideoListResponse struct {
	Videos []*models.Video `json:"video_list,omitempty"`
}

func QueryFavoriteVideoList(userId int32) (*FavoriteVideoListResponse,error){
	return NewFavoriteVideoListFlow(userId).Do()
}

func NewFavoriteVideoListFlow(userId int32) *FavoriteVideoListFlow{
	return &FavoriteVideoListFlow{
		UserId: userId,
	}
}

type FavoriteVideoListFlow struct {
	UserId int32

	*FavoriteVideoListResponse
	videos []*models.Video
}


func (f *FavoriteVideoListFlow) Do() (*FavoriteVideoListResponse,error){
	if err := f.Check();err != nil {
		return nil, err
	}

	if err := f.prepareData();err != nil{
		return nil,err
	}

	f.packResponse()

	return f.FavoriteVideoListResponse,nil
}

// Check 参数检查
func (f *FavoriteVideoListFlow) Check() error{
	if !models.NewUserInfoDAO().IsUserExitsById(f.UserId){
		return errors.New("user not found")
	}

	return nil
}

//数据处理
func (f *FavoriteVideoListFlow) prepareData() error{
	//mysql 链接表存储
	//dao := models.NewUserInfoDAO()
	//videos,err := dao.GetFavoriteVideoListByUserId(f.UserId)
	//if err != nil{
	//	return err
	//}
	//
	//f.videos = videos


	//redis存储
	dao := models.NewVideoDAO()
	redisOp := cache.NewRedisOperation()
	list,err := redisOp.GetFavoriteVideoIdList(f.UserId)
	if err != nil{
		return err
	}

	videos := make([]*models.Video,0,len(list))
	for _,rawId := range list{
		id,err := strconv.ParseInt(rawId,10,64)
		if err != nil{
			return err
		}
		video,err := dao.GetVideoByVideoId(int32(id))
		if err != nil{
			return err
		}
		videos = append(videos,video)
	}

	f.videos = videos

	return nil
}

//打包response
func (f *FavoriteVideoListFlow) packResponse(){
	f.FavoriteVideoListResponse = &FavoriteVideoListResponse{
		Videos: f.videos,
	}
}