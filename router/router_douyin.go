package router

import (
	"github.com/gin-gonic/gin"
	"github.com/hakusai22/douyin/controller/comment"
	"github.com/hakusai22/douyin/controller/user_info"
	"github.com/hakusai22/douyin/controller/user_login"
	"github.com/hakusai22/douyin/controller/video"
	"github.com/hakusai22/douyin/middlewares"
	"github.com/hakusai22/douyin/models"
)

func InitDouyinRouter() *gin.Engine {
	//初始化数据库表
	models.InitDB()

	// 获取Engine
	r := gin.Default()

	//设置静态文件夹 存储视频图片
	r.Static("static", "./static")

	//设置统一入口
	baseGroup := r.Group("/douyin")
	//根据灵活性考虑是否加入JWT中间件来进行鉴权，还是在之后再做鉴权
	// basic apis 基础api
	//视频推荐
	baseGroup.GET("/feed/", video.FeedVideoListHandler)
	//用户信息
	baseGroup.GET("/user/", middlewares.JWTMiddleWare(), user_info.UserInfoHandler)
	//登录
	baseGroup.POST("/user/login/", middlewares.SHAMiddleWare(), user_login.UserLoginHandler)
	//注册
	baseGroup.POST("/user/register/", middlewares.SHAMiddleWare(), user_login.UserRegisterHandler)
	//发布视频
	baseGroup.POST("/publish/action/", middlewares.JWTMiddleWare(), video.PublishVideoHandler)
	//视频列表
	baseGroup.GET("/publish/list/", middlewares.JWTMiddleWare(), video.QueryVideoListHandler)

	//extend 1 扩展1
	//进行点赞操作
	baseGroup.POST("/favorite/action/", middlewares.JWTMiddleWare(), video.PostFavorHandler)
	// 点赞列表
	baseGroup.GET("/favorite/list/", middlewares.JWTMiddleWare(), video.QueryFavorVideoListHandler)
	// 评论
	baseGroup.POST("/comment/action/", middlewares.JWTMiddleWare(), comment.PostCommentHandler)
	//评论列表
	baseGroup.GET("/comment/list/", middlewares.JWTMiddleWare(), comment.QueryCommentListHandler)

	//extend 2 扩展2
	//进行关注操作
	baseGroup.POST("/relation/action/", middlewares.JWTMiddleWare(), user_info.PostFollowActionHandler)
	//关注的列表
	baseGroup.GET("/relation/follow/list/", middlewares.JWTMiddleWare(), user_info.QueryFollowListHandler)
	//粉丝列表
	baseGroup.GET("/relation/follower/list/", middlewares.JWTMiddleWare(), user_info.QueryFollowerHandler)
	return r
}
