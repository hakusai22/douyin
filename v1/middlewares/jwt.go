package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	models2 "github.com/hakusai22/douyin/v1/models"
	"net/http"
	"time"
)

// jwt密钥
var jwtKey = []byte("0xCAFEBABY")

type Claims struct {
	UserId int64
	jwt.StandardClaims
}

// ReleaseToken 颁发token
func ReleaseToken(user models2.UserLogin) (string, error) {
	//7天有效时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.UserInfoId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "hakusai22",
			Subject:   "douyin",
		}}

	// HS256加密 claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析token 返回Claims
func ParseToken(tokenString string) (*Claims, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if token != nil {
		if key, ok := token.Claims.(*Claims); ok {
			if token.Valid {
				return key, true
			} else {
				return key, false
			}
		}
	}
	return nil, false
}

// JWTMiddleWare 直接拦截token 解析出userId
// JWTMiddleWare 鉴权中间件，鉴权并设置user_id
func JWTMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			tokenStr = c.PostForm("token")
		}
		//用户不存在
		if tokenStr == "" {
			c.JSON(http.StatusOK, models2.CommonResponse{StatusCode: 401, StatusMsg: "该用户不存在"})
			c.Abort() //阻止执行
			return
		}
		//验证token
		tokenStruck, ok := ParseToken(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, models2.CommonResponse{
				StatusCode: 403,
				StatusMsg:  "token不正确",
			})
			c.Abort() //阻止执行
			return
		}
		//token超时
		if time.Now().Unix() > tokenStruck.ExpiresAt {
			c.JSON(http.StatusOK, models2.CommonResponse{
				StatusCode: 402,
				StatusMsg:  "token过期",
			})
			c.Abort() //阻止执行
			return
		}
		c.Set("user_id", tokenStruck.UserId)
		c.Next()
	}
}
