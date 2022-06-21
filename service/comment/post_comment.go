package comment

import (
	"errors"
	"fmt"
	"github.com/hakusai22/douyin/models"
	"github.com/hakusai22/douyin/util"
)

const (
	CREATE = 1
	DELETE = 2
)

// Response 评论结构响应体
type Response struct {
	MyComment *models.Comment `json:"comment"`
}

// PostComment 封装数据 并调用do
func PostComment(userId int64, videoId int64, commentId int64, actionType int64, commentText string) (*Response, error) {
	return NewPostCommentFlow(userId, videoId, commentId, actionType, commentText).Do()
}

// PostCommentFlow 提交评论流结构体
type PostCommentFlow struct {
	userId      int64
	videoId     int64
	commentId   int64
	actionType  int64
	commentText string
	comment     *models.Comment
	*Response
}

// NewPostCommentFlow 封装
func NewPostCommentFlow(userId int64, videoId int64, commentId int64, actionType int64, commentText string) *PostCommentFlow {
	return &PostCommentFlow{userId: userId, videoId: videoId, commentId: commentId, actionType: actionType, commentText: commentText}
}

//Do 三部曲
func (p *PostCommentFlow) Do() (*Response, error) {
	var err error
	if err = p.checkNum(); err != nil {
		return nil, err
	}
	if err = p.prepareData(); err != nil {
		return nil, err
	}
	if err = p.packData(); err != nil {
		return nil, err
	}
	return p.Response, err
}

// CreateComment 增加评论
func (p *PostCommentFlow) CreateComment() (*models.Comment, error) {
	comment := models.Comment{UserInfoId: p.userId, VideoId: p.videoId, Content: p.commentText}
	err := models.NewCommentDAO().AddCommentAndUpdateCount(&comment)
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

// DeleteComment 删除评论
func (p *PostCommentFlow) DeleteComment() (*models.Comment, error) {
	//定义一个 comment
	var comment models.Comment
	err := models.NewCommentDAO().QueryCommentById(p.commentId, &comment)
	if err != nil {
		return nil, err
	}
	//删除comment
	err = models.NewCommentDAO().DeleteCommentAndUpdateCountById(p.commentId, p.videoId)
	if err != nil {
		return nil, err
	}
	//&引用 返回用户的评论
	return &comment, nil
}

//检查参数是否异常
func (p *PostCommentFlow) checkNum() error {
	if !models.NewUserInfoDAO().IsUserExistById(p.userId) {
		return fmt.Errorf("用户%d不存在", p.userId)
	}
	if !models.NewVideoDAO().IsVideoExistById(p.videoId) {
		return fmt.Errorf("视频%d不存在", p.videoId)
	}
	if p.actionType != CREATE && p.actionType != DELETE {
		return errors.New("未定义的行为")
	}
	return nil
}

//根据actionType调用不同的情况
func (p *PostCommentFlow) prepareData() error {
	var err error
	switch p.actionType {
	case CREATE:
		p.comment, err = p.CreateComment()
	case DELETE:
		p.comment, err = p.DeleteComment()
	default:
		return errors.New("未定义的操作")
	}
	return err
}

// 打包数据
func (p *PostCommentFlow) packData() error {
	//填充字段
	userInfo := models.UserInfo{}
	_ = models.NewUserInfoDAO().QueryUserInfoById(p.comment.UserInfoId, &userInfo)
	p.comment.User = userInfo
	_ = util.FillCommentFields(p.comment)

	p.Response = &Response{MyComment: p.comment}

	return nil
}
