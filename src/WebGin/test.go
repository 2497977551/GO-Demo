package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func PostParms(c *gin.Context) {
	form := &LoginForm{}

	if c.BindJSON(&form) == nil {
		fmt.Println(form.User, form.Password)
		if form.User == "user" && form.Password == "password" {
			c.JSON(200, gin.H{"status": "you are logged in"})
		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}
}
func main() {
	r := gin.Default()
	r.POST("/test", PostParms)
	err := r.Run(":1000")
	if err != nil {
		fmt.Println("服务启动失败：", err.Error())
	} else {
		fmt.Println("服务启动成功！")
	}
}
