package comment

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hakusai22/douyin/models"
	"github.com/hakusai22/douyin/service/comment"
	"net/http"
	"strconv"
)

// PostCommentResponse 提交评论响应体
type PostCommentResponse struct {
	models.CommonResponse
	*comment.Response
}

// PostCommentHandler handler
func PostCommentHandler(c *gin.Context) {
	NewProxyPostCommentHandler(c).Do()
}

type ProxyPostCommentHandler struct {
	*gin.Context
	videoId     int64
	userId      int64
	commentId   int64
	actionType  int64
	commentText string
}

func NewProxyPostCommentHandler(context *gin.Context) *ProxyPostCommentHandler {
	return &ProxyPostCommentHandler{Context: context}
}

func (p *ProxyPostCommentHandler) Do() {
	// 参数判断
	if err := p.parseNum(); err != nil {
		p.SendError(err.Error())
		return
	}
	// 调用service层的方法 进数据库
	commentRes, err := comment.PostComment(p.userId, p.videoId, p.commentId, p.actionType, p.commentText)
	if err != nil {
		p.SendError(err.Error())
		return
	}
	// 成功返回
	p.SendOk(commentRes)
}

func (p *ProxyPostCommentHandler) parseNum() error {
	//userId
	rawUserId, _ := p.Get("user_id")
	userId, ok := rawUserId.(int64)
	if !ok {
		return errors.New("userId解析出错")
	}
	p.userId = userId

	//视频id
	rawVideoId := p.Query("video_id")
	videoId, err := strconv.ParseInt(rawVideoId, 10, 64)
	if err != nil {
		return err
	}
	p.videoId = videoId

	//根据actionType解析对应的可选参数
	rawActionType := p.Query("action_type")
	actionType, err := strconv.ParseInt(rawActionType, 10, 64)
	switch actionType {
	//创建
	case comment.CREATE:
		p.commentText = p.Query("comment_text")
	//删除
	case comment.DELETE:
		p.commentId, err = strconv.ParseInt(p.Query("comment_id"), 10, 64)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("未定义的行为%d", actionType)
	}
	p.actionType = actionType
	return nil
}

// SendError error
func (p *ProxyPostCommentHandler) SendError(msg string) {
	p.JSON(http.StatusOK, PostCommentResponse{
		CommonResponse: models.CommonResponse{StatusCode: 1, StatusMsg: msg}, Response: &comment.Response{}})
}

// SendOk ok
func (p *ProxyPostCommentHandler) SendOk(comment *comment.Response) {
	p.JSON(http.StatusOK, PostCommentResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0},
		Response:       comment,
	})
}
