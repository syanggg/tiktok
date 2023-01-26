package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	ErrNilPtr        = errors.New("nil Ptr err")
	ErrEmptyUserList = errors.New("empty User list err")
	ErrUserNotFound  = errors.New("user is not found")
)

func NewUserInfo(name string, userLogin *UserLogin) *UserInfo {
	return &UserInfo{
		Name:      name,
		UserLogin: userLogin,
	}
}

type UserInfo struct {
	Id            int32       `json:"id" gorm:"id,primary_key"`
	Name          string      `json:"name" gorm:"name"`
	FollowCount   int32       `json:"follow_count" gorm:"follow_count"`
	FollowerCount int32       `json:"follower_count" gorm:"follower_count"`
	IsFollow      bool        `json:"is_follow" gorm:"is_follow"`
	UserLogin     *UserLogin  `json:"-"`
	Videos        []*Video    `json:"-"`
	Follows       []*UserInfo `json:"-" gorm:"many2many:user_relations;"`
	FavorVideos   []*Video    `json:"-" gorm:"many2many:user_favor_videos;"`
	Comments      []*Comment  `json:"-"`
}

// UserInfoDAO userinfo数据操作
type UserInfoDAO struct {
}

var (
	userInfoDAO  *UserInfoDAO
	userInfoOnce sync.Once //只能执行一次动作的对象
)

func NewUserInfoDAO() *UserInfoDAO {
	//保证只执行一次
	userInfoOnce.Do(func() {
		//userInfoDAO = &UserInfoDAO{}
		userInfoDAO = new(UserInfoDAO)
	})

	return userInfoDAO
}

// GetUserInfoById  通过用户id数据库表中查询用户
func (u *UserInfoDAO) GetUserInfoById(userId int32,)  (*UserInfo,error) {
	var info *UserInfo
	//select id,name,follow_count,follower_count,is_follow from userInfo where id=userId order by id limit 1;
	DB.Where("id = ?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(&info)

	//id = 0 表示查询失败，用户不存在
	if info.Id == 0 {
		return nil,ErrUserNotFound
	}

	return info,nil
}

// AddUserInfo 向数据库表中插入数据
func (u *UserInfoDAO) AddUserInfo(info *UserInfo) error {
	if info == nil {
		return ErrNilPtr
	}
	//insert into userInfo(...)  value(...)
	res := DB.Create(info)

	return res.Error
}

// RemoveUserInfo 移除当前用户信息
func (u *UserInfoDAO) RemoveUserInfo(info *UserInfo) error {
	if info == nil {
		return ErrNilPtr
	}

	//判断当前用户是否存在
	if !u.IsUserExitsById(info.Id) {
		return ErrUserNotFound
	}

	//delete from userInfo where id = info.Id;
	return DB.Delete(info).Error
}

// IsUserExitsById 判断用户是否存在
func (u *UserInfoDAO) IsUserExitsById(id int32) bool {
	var info UserInfo
	//select id from userInfo where id = id order by limit 1;
	//First获取查询结果的第一条数据
	if err := DB.Where("id = ?", id).Select("id").First(&info).Error; err != nil {
		log.Println(err)
	}

	if info.Id == 0 {
		return false
	}

	return true
}

// GetAllUserInfo 查询所有用户
//func (u *UserInfoDAO) GetAllUserInfo() ([]*UserInfo,error){
//	var allUserInfo	[]*UserInfo
//	if err := DB.Find(&allUserInfo).Error;err != nil{
//		return nil,err
//	}
//	return allUserInfo,nil
//}

// AddUserFollow 关注用户 更新数据库
func (u *UserInfoDAO) AddUserFollow(userId, userToId int32) error {
	//Transaction 事务
	return DB.Transaction(func(tx *gorm.DB) error {

		//update userInfo set follow_count = follow_count + 1 where id = userId;
		if err := DB.Model(&UserInfo{}).Where("id = ?", userId).Update("follow_count",
			gorm.Expr("follow_count + ?", 1)).Error; err != nil {
			return err
		}

		//update userInfo set follower_count = follower_count + 1 where id = userToId;
		if err := DB.Model(&UserInfo{}).Where("id = ?", userToId).Update("follower_count",
			gorm.Expr("follower_count + ?", 1)).Error; err != nil {
			return err
		}

		//获取当前用户
		info,err := u.GetUserInfoById(userId)
		if err != nil {
			return err
		}

		//获取被关注用户
		followerInfo,err := u.GetUserInfoById(userToId)
		if err != nil {
			return err
		}

		//关联表添加关联
		if err := DB.Model(info).Association("Follows").Append(followerInfo); err != nil {
			return err
		}

		return nil
	})
}

// CancelUserFollow 取消关注 更新数据库
func (u *UserInfoDAO) CancelUserFollow(userId, userToId int32) error {
	//事务
	return DB.Transaction(func(tx *gorm.DB) error {
		//获取当前用户
		info,err := u.GetUserInfoById(userId)
		if err != nil {
			return err
		}

		//获取被关注用户
		followerInfo,err := u.GetUserInfoById(userToId)
		if err != nil {
			return err
		}

		//update userInfo set follow_count = follow_count + 1 where id = info.Id
		if err := DB.Model(info).Update("follow_count", gorm.Expr("follow_count - ?", 1)).Error; err != nil {
			return err
		}

		//update userInfo set follower_count = follower_count + 1 where id = followInfo.Id
		if err := DB.Model(followerInfo).Update("follower_count", gorm.Expr("follower_count - ?", 1)).Error; err != nil {
			return err
		}

		//关联表删除关联
		if err := DB.Model(info).Association("Follows").Delete(followerInfo); err != nil {
			return err
		}

		return nil
	})
}

// GetFollowListById 通过用户id获取关注列表
func (u *UserInfoDAO) GetFollowListById(userId int32) ([]*UserInfo, error) {
	//获取当前用户
	info,err := u.GetUserInfoById(userId)
	if err != nil {
		return nil,err
	}

	var followInfos []*UserInfo
	//Find传入集合获取多条
	if err := DB.Model(info).Association("Follows").Find(&followInfos); err != nil {
		return nil, err
	}

	return followInfos, nil
}

// GetFollowerListById 通过用户id获取粉丝列表
func (u *UserInfoDAO) GetFollowerListById(userId int32) ([]*UserInfo, error) {
	var followerInfos []*UserInfo

	//查询链接表中关注userId用户的用户
	if err := DB.Raw("select * from user_infos where id in (select user_info_id from user_relations where follow_id = ?)",
		userId).Scan(&followerInfos).Error; err != nil {
		return nil, err
	}

	return followerInfos, nil
}

// GetFavoriteVideoListByUserId 通过用户id获取喜欢的视频列表
func (u *UserInfoDAO) GetFavoriteVideoListByUserId(id int32) ([]*Video,error){
	var favoriteVideos []*Video

	//当前用户
	info,err := u.GetUserInfoById(id)
	if err != nil {
		return nil,err
	}

	//查询链接表中的数据
	if err := DB.Model(info).Association("FavorVideos").Find(&favoriteVideos);err != nil{
		return nil,err
	}

	return favoriteVideos,nil
}