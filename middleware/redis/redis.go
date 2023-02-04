package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"tiktok/util"
)

/*
	redis存储 后面可以将关注列表 mysql移到redis
 */

var rdb *redis.Client
var c  = context.Background()

func InitRedis() error{
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",util.Info.RDB.Host,util.Info.RDB.Port),
		Password: "",
		DB: util.Info.RDB.Database,
		PoolSize: 100,
	})

	_,err := rdb.Ping(c).Result()
	if err != nil{
		return err
	}

	return nil
}

type RedisOperation struct {
}

func NewRedisOperation() *RedisOperation{
	return &RedisOperation{}
}

// AddFavorite 添加喜欢
func (r *RedisOperation) AddFavorite(userId int32,videoId int32) error{
	key := fmt.Sprintf("favorite:%d",userId)
	//将videoId添加到key为favorite:userId的set集合
	err := rdb.SAdd(c,key,videoId).Err()
	if err != nil{
		return err
	}

	return nil
}

// CancelFavorite 取消喜欢
func (r *RedisOperation) CancelFavorite(userId int32,videoId int32) error{
	key := fmt.Sprintf("favorite:%d",userId)
	//将videoId从key为favorite:userId的set集合移除
	err := rdb.SRem(c,key,videoId).Err()
	if err != nil{
		return err
	}

	return nil
}

// GetFavoriteState 获取点赞状态
func (r *RedisOperation) GetFavoriteState(userId int32,videoId int32) (bool,error){
	key := fmt.Sprintf("favorite:%d",userId)
	//判断videoId是否时key集合的成员
	result,err := rdb.SIsMember(c,key,videoId).Result()
	if err != nil{
		return result,err
	}

	return result,nil
}

// GetFavoriteVideoIdList 获取点赞的视频Id列表
func (r *RedisOperation) GetFavoriteVideoIdList(userId int32) ([]string,error){
	key := fmt.Sprintf("favorite:%d",userId)

	list,err := rdb.SMembers(c,key).Result()

	return list,err
}

// AddFollow 添加关注
func (r *RedisOperation) AddFollow(userId int32,followId int32) error{
	key := fmt.Sprintf("follow:%d",userId)

	err := rdb.SAdd(c,key,followId).Err()
	if err != nil{
		return err
	}

	return nil
}

// CanCelFollow 取消关注
func (r *RedisOperation) CanCelFollow(userId int32,followId int32) error{
	key := fmt.Sprintf("follow:%d",userId)

	err := rdb.SRem(c,key,followId).Err()
	if err != nil{
		return err
	}

	return nil
}

// GetFollowState 获取关注状态
func (r *RedisOperation) GetFollowState(userId int32,followId int32) (bool,error){
	key := fmt.Sprintf("follow:%d",userId)

	result,err := rdb.SIsMember(c,key,followId).Result()
	if err != nil{
		return result,err
	}

	return result,err
}