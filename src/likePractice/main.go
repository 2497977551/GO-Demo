package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	var err error
	err, Db = IndexInit()
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	} else {
		fmt.Println("连接成功")
	}
}
func main() {
	defer IndexDefer(Db)

	r := gin.Default()
	userGroup := r.Group("/user")
	userGroup.Use(Cors())
	{
		userGroup.POST("/register", RegisterUsers)
		userGroup.POST("/login", MemberLogin)
		userGroup.POST("/updateMember", UpdateMembers)
		userGroup.POST("/memberPostBlog", PostBlog)
	}

	if err := r.Run(":5050"); err != nil {
		fmt.Println("服务启动成功")
	}

}
