package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
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
	defer func() {
		db := IndexDefer(Db)
		err := db.Close()
		if err != nil {
			log.Fatalln("数据库连接断开失败", err.Error())
		} else {

			fmt.Println("数据库连接断开成功")
		}
		e := GetRecover()
		fmt.Println(e)
	}()

	r := gin.Default()
	userGroup := r.Group("/user")
	userGroup.Use(Cors())
	{
		userGroup.POST("/register", RegisterUsers)
		userGroup.POST("/login", MemberLogin)
		userGroup.POST("/updateMember", UpdateMembers)
		userGroup.POST("/memberPostBlog", PostBlog)
		userGroup.POST("/memberPostNotice", MemberPostNotice)
	}

	if err := r.Run(":5050"); err != nil {
		fmt.Println("服务启动成功")
	} else {
		fmt.Println("服务启动失败:", err)
	}

}
