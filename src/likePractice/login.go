package main

import (
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
	data := make(map[string]interface{})
	if err := c.ShouldBindJSON(&form); err != nil {
		panic(err.Error())
	} else {
		r, u := userLogin(form.Username, form.Password)
		if r.Error != nil {
			panic(r.Error)
		}
		if form.Username == u.Username && form.Password == u.Password {
			data["code"] = 200
			data["message"] = "登录成功！"
			c.JSON(http.StatusOK, gin.H{
				"data": data,
			})
		} else if u.Username == "" || u.Password == "" {
			data["code"] = 200
			data["message"] = "账号或者密码错误！"
			c.JSON(http.StatusOK, gin.H{
				"data": data,
			})

		} else {
			data["code"] = 401
			data["message"] = "登录失败！"
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": data,
			})
		}
	}
}
