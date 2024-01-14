package video

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hakusai22/douyin/v1/models"
	"github.com/hakusai22/douyin/v1/service/video"
	"net/http"
)

// FavorVideoListResponse 点赞视频列表
type FavorVideoListResponse struct {
	models.CommonResponse
	*video.FavorList
}

// QueryFavorVideoListHandler DO
func QueryFavorVideoListHandler(c *gin.Context) {
	NewProxyFavorVideoListHandler(c).Do()
}

// ProxyFavorVideoListHandler userId+context
type ProxyFavorVideoListHandler struct {
	*gin.Context
	userId int64
}

// NewProxyFavorVideoListHandler 封装一层
func NewProxyFavorVideoListHandler(c *gin.Context) *ProxyFavorVideoListHandler {
	return &ProxyFavorVideoListHandler{Context: c}
}

func (p *ProxyFavorVideoListHandler) Do() {
	//解析参数
	if err := p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	//调用Service层
	favorVideoList, err := video.QueryFavorVideoList(p.userId)
	if err != nil {
		p.SendError(err.Error())
		return
	}
	//成功返回json到gin
	p.SendOk(favorVideoList)
}

func (p *ProxyFavorVideoListHandler) parseNum() error {
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId
	return nil
}

func (p *ProxyFavorVideoListHandler) SendError(msg string) {
	p.JSON(http.StatusOK, FavorVideoListResponse{
		CommonResponse: models.CommonResponse{StatusCode: 1, StatusMsg: msg}})
}

func (p *ProxyFavorVideoListHandler) SendOk(favorList *video.FavorList) {
	p.JSON(http.StatusOK, FavorVideoListResponse{CommonResponse: models.CommonResponse{StatusCode: 0},
		FavorList: favorList,
	})
}
