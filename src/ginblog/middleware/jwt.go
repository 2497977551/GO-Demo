package middleware

import (
	"ginblog/utils"
	"ginblog/utils/ErrorInfo"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var JwtKey = []byte(utils.JwtKey)
var code int

type MyClaims struct {
	UserName string `json:"UserName"`

	jwt.StandardClaims
}

// 生成token
func SetToken(username string) (string, int) {
	expiresTime := time.Now().Add(24 * time.Hour) // 24小时有效时间
	setClaims := MyClaims{
		UserName: username,

		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresTime.Unix(),
			Issuer:    "joshua",
		},
	}
	// 生成token
	reqClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, setClaims)
	token, err := reqClaims.SignedString(JwtKey)
	if err != nil {
		return "", ErrorInfo.Error
	}
	return token, ErrorInfo.SucCse
}

// 验证token
func CheckToken(token string) (*MyClaims, int) {
	setToken, _ := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})
	if key, code := setToken.Claims.(*MyClaims); code && setToken.Valid {
		return key, ErrorInfo.SucCse
	} else {
		return nil, ErrorInfo.Error
	}
}

// JWT中间件
func JWTTokenMid() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.Request.Header.Get("Authorization")
		if tokenHeader == "" {
			code = ErrorInfo.ERRTokenNoExistent
			c.JSON(http.StatusOK, gin.H{
				"Status":  code,
				"Message": ErrorInfo.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		checkToken := strings.SplitN(tokenHeader, " ", 2)
		if len(checkToken) != 2 && checkToken[0] != "Bearer" {
			code = ErrorInfo.ERRTokenFormatWrong
			c.JSON(http.StatusOK, gin.H{
				"Status":  code,
				"Message": ErrorInfo.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		key, TCode := CheckToken(checkToken[1])
		if TCode == ErrorInfo.Error {
			code = ErrorInfo.ERRTokenWrong
			c.JSON(http.StatusOK, gin.H{
				"Status":  code,
				"Message": ErrorInfo.GetErrMsg(code),
			})
			c.Abort()
			return
		}
		if time.Now().Unix() > key.ExpiresAt {
			code = ErrorInfo.ERRTokenOverdue
			c.JSON(http.StatusOK, gin.H{
				"Status":  code,
				"Message": ErrorInfo.GetErrMsg(code),
			})
			c.Abort()
			return
		}

		c.Set("UserName", key.UserName)
		c.Next()
	}
}
