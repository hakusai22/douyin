package comment

import (
	"errors"
	"fmt"
	"github.com/hakusai22/douyin/models"
	"github.com/hakusai22/douyin/util"
)

// List list封装
type List struct {
	Comments []*models.Comment `json:"comment_list"`
}

// QueryCommentList service层的方法
func QueryCommentList(userId, videoId int64) (*List, error) {
	return NewQueryCommentListFlow(userId, videoId).Do()
}

// QueryCommentListFlow 评论列表
type QueryCommentListFlow struct {
	userId      int64
	videoId     int64
	comments    []*models.Comment //评论集合
	commentList *List
}

// NewQueryCommentListFlow 对userid  videoid 封装一层
func NewQueryCommentListFlow(userId, videoId int64) *QueryCommentListFlow {
	return &QueryCommentListFlow{userId: userId, videoId: videoId}
}

// Do 3步曲
func (q *QueryCommentListFlow) Do() (*List, error) {
	if err := q.checkNum(); err != nil {
		return nil, err
	}
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.commentList, nil
}

//检查参数
func (q *QueryCommentListFlow) checkNum() error {
	if !models.NewUserInfoDAO().IsUserExistById(q.userId) {
		return fmt.Errorf("用户%d处于登出状态", q.userId)
	}
	if !models.NewVideoDAO().IsVideoExistById(q.videoId) {
		return fmt.Errorf("视频%d不存在或已经被删除", q.videoId)
	}
	return nil
}

//数据库进行查询
func (q *QueryCommentListFlow) prepareData() error {
	err := models.NewCommentDAO().QueryCommentListByVideoId(q.videoId, &q.comments)
	if err != nil {
		return err
	}
	//根据前端的要求填充正确的时间格式
	err = util.FillCommentListFields(&q.comments)
	if err != nil {
		return errors.New("暂时还没有人评论")
	}
	return nil
}

//打包 封装
func (q *QueryCommentListFlow) packData() error {
	//评论集合 将q里面的comments封装成List形式
	q.commentList = &List{Comments: q.comments}
	return nil
}
