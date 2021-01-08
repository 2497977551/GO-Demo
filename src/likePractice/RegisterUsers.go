package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// 创建数据函数
func registerUser(username, password string, creationTime, updateTime time.Time) (res *gorm.DB) {
	user := Member{
		Username:     username,
		Password:     password,
		CreationTime: creationTime,
		UpdateTime:   updateTime,
	}
	res = Db.Create(&user)
	return
}
func RegisterUsers(c *gin.Context) {
	u := Member{}

	if err := c.ShouldBindJSON(&u); err != nil {
		fmt.Println(err.Error())
		Data["code"] = 401
		Data["message"] = "类型错误！"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": Data,
		})
	} else {
		if u.Username != "" && u.Password != "" {
			res := registerUser(u.Username, u.Password, time.Now(), time.Now())
			if res.Error != nil {
				panic(res.Error)
			}
			Data["code"] = 200
			Data["message"] = "注册成功!"
			c.JSON(http.StatusOK, gin.H{
				"data": Data,
			})
		} else {
			Data["code"] = 401
			Data["message"] = "注册失败!"
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": Data,
			})
		}
	}
}
