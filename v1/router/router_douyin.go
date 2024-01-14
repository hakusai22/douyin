package router

import (
	comment2 "douyin/v1/controller/comment"
	user_info2 "douyin/v1/controller/user_info"
	user_login2 "douyin/v1/controller/user_login"
	video2 "douyin/v1/controller/video"
	middlewares2 "douyin/v1/middlewares"
	"douyin/v1/models"
	"github.com/gin-gonic/gin"
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
	baseGroup.GET("/feed/", video2.FeedVideoListHandler)
	//用户信息
	baseGroup.GET("/user/", middlewares2.JWTMiddleWare(), user_info2.UserInfoHandler)
	//登录
	baseGroup.POST("/user/login/", middlewares2.SHAMiddleWare(), user_login2.UserLoginHandler)
	//注册
	baseGroup.POST("/user/register/", middlewares2.SHAMiddleWare(), user_login2.UserRegisterHandler)
	//发布视频
	baseGroup.POST("/publish/action/", middlewares2.JWTMiddleWare(), video2.PublishVideoHandler)
	//视频列表
	baseGroup.GET("/publish/list/", middlewares2.JWTMiddleWare(), video2.QueryVideoListHandler)

	//extend 1 扩展1
	//进行点赞操作
	baseGroup.POST("/favorite/action/", middlewares2.JWTMiddleWare(), video2.PostFavorHandler)
	// 点赞列表
	baseGroup.GET("/favorite/list/", middlewares2.JWTMiddleWare(), video2.QueryFavorVideoListHandler)
	// 评论
	baseGroup.POST("/comment/action/", middlewares2.JWTMiddleWare(), comment2.PostCommentHandler)
	//评论列表
	baseGroup.GET("/comment/list/", middlewares2.JWTMiddleWare(), comment2.QueryCommentListHandler)

	//extend 2 扩展2
	//进行关注操作
	baseGroup.POST("/relation/action/", middlewares2.JWTMiddleWare(), user_info2.PostFollowActionHandler)
	//关注的列表
	baseGroup.GET("/relation/follow/list/", middlewares2.JWTMiddleWare(), user_info2.QueryFollowListHandler)
	//粉丝列表
	baseGroup.GET("/relation/follower/list/", middlewares2.JWTMiddleWare(), user_info2.QueryFollowerHandler)
	return r
}
