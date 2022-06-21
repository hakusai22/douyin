package models

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"sync"
)

//全局错误
var (
	ErrIvdPtr        = errors.New("空指针错误")
	ErrEmptyUserList = errors.New("用户列表为空")
)

// UserInfo 用户信息结构体
type UserInfo struct {
	Id            int64       `json:"id" gorm:"id,omitempty"`
	Name          string      `json:"name" gorm:"name,omitempty"`
	FollowCount   int64       `json:"follow_count" gorm:"follow_count,omitempty"`
	FollowerCount int64       `json:"follower_count" gorm:"follower_count,omitempty"`
	IsFollow      bool        `json:"is_follow" gorm:"is_follow,omitempty"`
	User          *UserLogin  `json:"-"`                                     //用户与密码之间的多对多
	Videos        []*Video    `json:"-"`                                     //用户与投稿视频的一对多
	Follows       []*UserInfo `json:"-" gorm:"many2many:user_relations;"`    //用户之间的多对多
	FavorVideos   []*Video    `json:"-" gorm:"many2many:user_favor_videos;"` //用户与点赞视频之间的多对多
	Comments      []*Comment  `json:"-"`                                     //用户与评论的一对多
}

// UserInfoDAO 用户dao
type UserInfoDAO struct {
}

// 全局dao 单例 (它能够让函数方法只执行一次)
var (
	userInfoDAO  *UserInfoDAO
	userInfoOnce sync.Once
)

// NewUserInfoDAO 创建用户dao函数
func NewUserInfoDAO() *UserInfoDAO {
	userInfoOnce.Do(func() {
		userInfoDAO = new(UserInfoDAO)
	})
	return userInfoDAO
}

// QueryUserInfoById 查询用户信息通过id
func (u *UserInfoDAO) QueryUserInfoById(userId int64, userinfo *UserInfo) error {
	if userinfo == nil {
		return ErrIvdPtr
	}
	//DB.Where("id=?",userId).First(userinfo)
	DB.Where("id=?", userId).Select([]string{"id", "name", "follow_count", "follower_count", "is_follow"}).First(userinfo)
	//id为零值，说明sql执行失败
	if userinfo.Id == 0 {
		return errors.New("该用户不存在")
	}
	return nil
}

// AddUserInfo 添加用户并返回用户信息
func (u *UserInfoDAO) AddUserInfo(userinfo *UserInfo) error {
	if userinfo == nil {
		return ErrIvdPtr
	}
	return DB.Create(userinfo).Error
}

// IsUserExistById 判断用户是否存在
func (u *UserInfoDAO) IsUserExistById(id int64) bool {
	var userinfo UserInfo
	if err := DB.Where("id=?", id).Select("id").First(&userinfo).Error; err != nil {
		log.Println(err)
	}
	if userinfo.Id == 0 {
		return false
	}
	return true
}

// AddUserFollow 添加用户关注
func (u *UserInfoDAO) AddUserFollow(userId, userToId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE user_infos SET follow_count=follow_count+1 WHERE id = ?", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_infos SET follower_count=follower_count+1 WHERE id = ?", userToId).Error; err != nil {
			return err
		}
		if err := tx.Exec("INSERT INTO `user_relations` (`user_info_id`,`follow_id`) VALUES (?,?)", userId, userToId).Error; err != nil {
			return err
		}
		return nil
	})
}

// CancelUserFollow 取消用户关注
func (u *UserInfoDAO) CancelUserFollow(userId, userToId int64) error {
	return DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("UPDATE user_infos SET follow_count=follow_count-1 WHERE id = ? AND follow_count>0", userId).Error; err != nil {
			return err
		}
		if err := tx.Exec("UPDATE user_infos SET follower_count=follower_count-1 WHERE id = ? AND follower_count>0", userToId).Error; err != nil {
			return err
		}
		if err := tx.Exec("DELETE FROM `user_relations` WHERE user_info_id=? AND follow_id=?", userId, userToId).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetFollowListByUserId 通过用户id获取到关注列表
func (u *UserInfoDAO) GetFollowListByUserId(userId int64, userList *[]*UserInfo) error {
	if userList == nil {
		return ErrIvdPtr
	}
	var err error
	if err = DB.Raw("SELECT u.* FROM user_relations r, user_infos u WHERE r.user_info_id = ? AND r.follow_id = u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	if len(*userList) == 0 || (*userList)[0].Id == 0 {
		return ErrEmptyUserList
	}
	return nil
}

// GetFollowerListByUserId 按用户 ID 获取粉丝列表
func (u *UserInfoDAO) GetFollowerListByUserId(userId int64, userList *[]*UserInfo) error {
	if userList == nil {
		return ErrIvdPtr
	}
	var err error
	if err = DB.Raw("SELECT u.* FROM user_relations r, user_infos u WHERE r.follow_id = ? AND r.user_info_id = u.id", userId).Scan(userList).Error; err != nil {
		return err
	}
	//if len(*userList) == 0 || (*userList)[0].Id == 0 {
	//	return ErrEmptyUserList
	//}
	return nil
}
