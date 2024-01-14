package models

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

// Comment 评论结构体
type Comment struct {
	Id         int64     `json:"id"`
	UserInfoId int64     `json:"-"` //用于一对多关系的id
	VideoId    int64     `json:"-"` //一对多，视频对评论
	User       UserInfo  `json:"user" gorm:"-"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"-"`
	CreateDate string    `json:"create_date" gorm:"-"`
}

// CommentDAO dao结构体
type CommentDAO struct {
}

// 全局定义CommentDAO
var (
	commentDao CommentDAO
)

// NewCommentDAO 创建dao
func NewCommentDAO() *CommentDAO {
	return &commentDao
}

// AddCommentAndUpdateCount 增加评论更新数量
func (c *CommentDAO) AddCommentAndUpdateCount(comment *Comment) error {
	if comment == nil {
		return errors.New("AddCommentAndUpdateCount comment空指针")
	}
	//执行事务
	return DB.Transaction(func(tx *gorm.DB) error {
		//添加评论数据
		if err := tx.Create(comment).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		//增加count
		if err := tx.Exec("UPDATE videos v SET v.comment_count = v.comment_count+1 WHERE v.id=?", comment.VideoId).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

// DeleteCommentAndUpdateCountById 删除评论更新数量
func (c *CommentDAO) DeleteCommentAndUpdateCountById(commentId, videoId int64) error {
	//执行事务
	return DB.Transaction(func(tx *gorm.DB) error {
		//删除评论
		if err := tx.Exec("DELETE FROM comments WHERE id = ?", commentId).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		//减少count
		if err := tx.Exec("UPDATE videos v SET v.comment_count = v.comment_count-1 WHERE v.id=? AND v.comment_count>0", videoId).Error; err != nil {
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

// QueryCommentById 查询评论通过用户id
func (c *CommentDAO) QueryCommentById(id int64, comment *Comment) error {
	if comment == nil {
		return errors.New("QueryCommentById comment 空指针")
	}
	// 查出 给 comment *Comment赋值
	return DB.Where("id=?", id).First(comment).Error
}

// QueryCommentListByVideoId 查询评论列表通过视频id    comments *[]*Comment 传入一个空的对面进行 进行赋值
func (c *CommentDAO) QueryCommentListByVideoId(videoId int64, comments *[]*Comment) error {
	if comments == nil {
		return errors.New("QueryCommentListByVideoId comments空指针")
	}
	//查出赋值
	if err := DB.Model(&Comment{}).Where("video_id=?", videoId).Find(comments).Error; err != nil {
		return err
	}
	return nil
}
