package v1

import (
	"fmt"
	"ginblog/middleware"
	"ginblog/model"
	"ginblog/utils/ErrorInfo"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
	//"reflect"
	"strconv"
	"time"
)

var (
	code int
	err  error
)

// 查询用户是否存在
func QueryUserIfExist(c *gin.Context) {
	fmt.Println("查询用户是否存在")
}

// 添加用户
func AddUser(c *gin.Context) {
	from := model.User{CreationTime: time.Now(), Id: uuid.NewV1()}
	err = c.ShouldBindJSON(&from)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusNotFound, gin.H{
				"Status":  404,
				"MessAge": "参数错误",
			})
			fmt.Println(err)
		} else {
			code = model.CheckUser(from.UserName)
			if code == ErrorInfo.SucCse {
				code = model.CreateUser(&from)
				if code == ErrorInfo.SucCse {
					c.JSON(http.StatusOK, gin.H{
						"Status":  code,
						"MessAge": ErrorInfo.GetErrMsg(code),
					})
				} else {
					c.JSON(http.StatusNotFound, gin.H{
						"Status":  code,
						"MessAge": ErrorInfo.GetErrMsg(code),
					})
				}

			} else {
				c.JSON(http.StatusNotFound, gin.H{
					"Status":  code,
					"MessAge": ErrorInfo.GetErrMsg(code),
				})
			}
		}
	}()

	panic(err)
}

// 查询单个用户
func QueryUser(c *gin.Context) {
	userName := c.Query("UserName")
	user, code := model.QueryUsers(userName)
	if code == ErrorInfo.SucCse {
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Data":    user,
			"MessAge": ErrorInfo.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"MessAge": ErrorInfo.GetErrMsg(code),
		})
	}
}

// 查询所有用户列表
func QueryAllUserList(c *gin.Context) {
	pageNumber, err1 := strconv.Atoi(c.Query("PageNumber"))
	pageSize, err2 := strconv.Atoi(c.Query("PageSize"))
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  404,
			"MessAge": "格式转换失败",
		})
	}
	if pageNumber > 0 && pageSize >= 20 {

		user, count, pageNumbers, pageSizes, code := model.GetAllUser(pageNumber, pageSize)
		if code == ErrorInfo.SucCse {
			data := map[string]interface{}{
				"UserInfo":   user,
				"Total":      count,
				"PageNumber": pageNumbers,
				"PageSize":   pageSizes,
			}
			c.JSON(http.StatusOK, gin.H{
				"Status":  code,
				"Data":    data,
				"MessAge": ErrorInfo.GetErrMsg(code),
			})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"Status":  code,
				"MessAge": ErrorInfo.GetErrMsg(code),
			})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"Status":  404,
			"MessAge": "请检查参数是否正确",
		})
	}

}

// 编辑用户

func EditUser(c *gin.Context) {

	from := model.Users{}
	err = c.ShouldBind(&from)
	if err != nil {
		log.Fatalln(err)
	}
	userCode := model.CheckUser(from.UserName)
	if userCode == 200 {
		code := model.UpdateUser(from.Id, from)
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": ErrorInfo.GetErrMsg(code),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"Status":  userCode,
			"Message": ErrorInfo.GetErrMsg(userCode),
		})
	}

}

// 删除用户
func DeleteUser(c *gin.Context) {
	from := model.User{}
	err = c.ShouldBindJSON(&from)
	code := model.DeleteUsers(from.Id)
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"Message": ErrorInfo.GetErrMsg(code),
	})
}

// 登录
func UserLogin(c *gin.Context) {
	from := model.UserLogin{}
	var (
		token string
		code  int
	)
	err = c.ShouldBindJSON(&from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status":  404,
			"Message": err,
		})
		return
	}
	code = model.Login(from.UserName, from.PassWord)

	if code == ErrorInfo.SucCse {
		token, code = middleware.SetToken(from.UserName)

	}
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"MessAge": ErrorInfo.GetErrMsg(code),
		"Token":   token,
	})
}
