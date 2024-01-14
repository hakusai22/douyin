package video

import (
	"douyin/v1/models"
	"douyin/v1/service/video"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListResponse struct {
	models.CommonResponse
	*video.List
}

func QueryVideoListHandler(c *gin.Context) {
	p := NewProxyQueryVideoList(c)
	rawId, _ := c.Get("user_id")
	err := p.DoQueryVideoListByUserId(rawId)
	if err != nil {
		p.QueryVideoListError(err.Error())
	}
}

// ProxyQueryVideoList  context代理类
type ProxyQueryVideoList struct {
	c *gin.Context
}

// NewProxyQueryVideoList 封装一层
func NewProxyQueryVideoList(c *gin.Context) *ProxyQueryVideoList {
	return &ProxyQueryVideoList{c: c}
}

// DoQueryVideoListByUserId 根据userId字段进行查询
func (p *ProxyQueryVideoList) DoQueryVideoListByUserId(rawId interface{}) error {
	userId, ok := rawId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	videoList, err := video.QueryVideoListByUserId(userId)
	if err != nil {
		return err
	}
	//json 返回到gin
	p.QueryVideoListOk(videoList)
	return nil
}

func (p *ProxyQueryVideoList) QueryVideoListError(msg string) {
	p.c.JSON(http.StatusOK, ListResponse{CommonResponse: models.CommonResponse{
		StatusCode: 1,
		StatusMsg:  msg,
	}})
}

func (p *ProxyQueryVideoList) QueryVideoListOk(videoList *video.List) {
	p.c.JSON(http.StatusOK, ListResponse{
		CommonResponse: models.CommonResponse{
			StatusCode: 0,
		},
		List: videoList,
	})
}
