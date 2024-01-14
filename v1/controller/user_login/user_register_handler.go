package user_login

import (
	"douyin/v1/models"
	user_login2 "douyin/v1/service/user_login"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegisterResponse 用户注册response 用户id+token
type UserRegisterResponse struct {
	models.CommonResponse
	*user_login2.LoginResponse
}

func UserRegisterHandler(c *gin.Context) {
	username := c.Query("username")
	rawVal, _ := c.Get("password")
	password, ok := rawVal.(string)
	if !ok {
		c.JSON(http.StatusOK, UserRegisterResponse{
			CommonResponse: models.CommonResponse{
				StatusCode: 1,
				StatusMsg:  "密码解析出错",
			},
		})
		return
	}
	//调用service方法
	registerResponse, err := user_login2.PostUserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, UserRegisterResponse{
			CommonResponse: models.CommonResponse{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	//josn返回数据到context里面去
	c.JSON(http.StatusOK, UserRegisterResponse{
		CommonResponse: models.CommonResponse{StatusCode: 0},
		LoginResponse:  registerResponse,
	})
}
