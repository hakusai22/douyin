package video

import (
	"errors"
	models2 "github.com/hakusai22/douyin/v1/models"
)

// FavorList 返回的结构
type FavorList struct {
	Videos []*models2.Video `json:"video_list"`
}

// QueryFavorVideoList handler 调用方法
func QueryFavorVideoList(userId int64) (*FavorList, error) {
	return NewQueryFavorVideoListFlow(userId).Do()
}

// QueryFavorVideoListFlow 封装的结构
type QueryFavorVideoListFlow struct {
	userId    int64
	videos    []*models2.Video
	videoList *FavorList
}

// NewQueryFavorVideoListFlow 封装一层
func NewQueryFavorVideoListFlow(userId int64) *QueryFavorVideoListFlow {
	return &QueryFavorVideoListFlow{userId: userId}
}

// Do 三部曲
func (q *QueryFavorVideoListFlow) Do() (*FavorList, error) {
	if err := q.checkNum(); err != nil {
		return nil, err
	}
	if err := q.prepareData(); err != nil {
		return nil, err
	}
	if err := q.packData(); err != nil {
		return nil, err
	}
	return q.videoList, nil
}

func (q *QueryFavorVideoListFlow) checkNum() error {
	if !models2.NewUserInfoDAO().IsUserExistById(q.userId) {
		return errors.New("用户状态异常")
	}
	return nil
}

func (q *QueryFavorVideoListFlow) prepareData() error {
	err := models2.NewVideoDAO().QueryFavorVideoListByUserId(q.userId, &q.videos)
	if err != nil {
		return err
	}
	//填充信息(Author和IsFavorite字段，由于是点赞列表，故所有的都是点赞状态
	for i := range q.videos {
		//作者信息查询
		var userInfo models2.UserInfo
		err = models2.NewUserInfoDAO().QueryUserInfoById(q.videos[i].UserInfoId, &userInfo)
		if err == nil { //若查询未出错则更新，否则不更新作者信息
			q.videos[i].Author = userInfo
		}
		q.videos[i].IsFavorite = true
	}
	return nil
}

func (q *QueryFavorVideoListFlow) packData() error {
	q.videoList = &FavorList{Videos: q.videos}
	return nil
}
