package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func userLogin(username, password string) (res *gorm.DB, user Member) {
	user = Member{}
	res = Db.Where("username = ? AND password = ?", username, password).First(&user)
	return
}

func MemberLogin(c *gin.Context) {
	form := Member{}
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println(err.Error())
		Data["code"] = 401
		Data["message"] = "类型错误！"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": Data,
		})
	} else {
		r, u := userLogin(form.Username, form.Password)
		if r.Error != nil {
			panic(r.Error)
		}
		if form.Username == u.Username && form.Password == u.Password {
			Data["code"] = 200
			Data["message"] = "登录成功！"
			c.JSON(http.StatusOK, gin.H{
				"data": Data,
			})
		} else if u.Username == "" || u.Password == "" {
			Data["code"] = 200
			Data["message"] = "账号或者密码错误！"
			c.JSON(http.StatusOK, gin.H{
				"data": Data,
			})

		} else {
			Data["code"] = 401
			Data["message"] = "登录失败！"
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": Data,
			})
		}
	}
}
