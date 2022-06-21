package user_login

import (
	"github.com/gin-gonic/gin"
	"github.com/hakusai22/douyin/models"
	"github.com/hakusai22/douyin/service/user_login"
	"net/http"
)

type UserLoginResponse struct {
	models.CommonResponse
	*user_login.LoginResponse
}

func UserLoginHandler(c *gin.Context) {
	username := c.Query("username")
	//从context中获取加密的密码
	raw, _ := c.Get("password")
	//判断是否是字符串
	password, ok := raw.(string)
	//ok=false
	if !ok {
		c.JSON(http.StatusOK, UserLoginResponse{
			CommonResponse: models.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "密码解析错误",
			},
		})
	}
	// 调用service 方法 查询信息
	userLoginResponse, err := user_login.QueryUserLogin(username, password)
	//用户存在返回对应的错误
	if err != nil {
		c.JSON(http.StatusOK, UserLoginResponse{
			CommonResponse: models.CommonResponse{StatusCode: 1, StatusMsg: err.Error()},
		})
		return
	}

	//用户存在，返回相应的id和token
	c.JSON(http.StatusOK, UserLoginResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0},
		LoginResponse:  userLoginResponse,
	})
}
