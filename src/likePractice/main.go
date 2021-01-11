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
	apiGroup := r.Group("/api")
	apiGroup.Use(Cors())
	{
		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("/register", RegisterUsers)          // 用户注册
			userGroup.POST("/login", MemberLogin)               // 登录
			userGroup.POST("/updateMember", UpdateMembers)      // 修改信息
			userGroup.POST("/queryMemberInfo", QueryMemberInfo) // 查询个人信息
			userGroup.GET("/queryAllFollow", QueryAllFollow)    // 查询多少人关注我
		}

		blogGroup := apiGroup.Group("/blog")
		{
			blogGroup.POST("/memberPostBlog", PostBlog)           // 发布博客
			blogGroup.POST("/memberPostNotice", MemberPostNotice) // 发布公告
			blogGroup.POST("/commentBlog", CommentBlog)           // 评论博客
			blogGroup.POST("/queryMaxBlog", QueryMaxBlog)         // 查询首页所有博客
			blogGroup.POST("/followMember", FollowMember)         // 关注或者取消关注博主
			blogGroup.POST("/queryAllBlog", QueryAllBlog)         // 查询已关注的用户所有博客
			blogGroup.POST("/memberStar", Star)                   // 点赞
		}

	}

	if err := r.Run(":5050"); err != nil {
		fmt.Println("服务启动成功")
	} else {
		fmt.Println("服务启动失败:", err)
	}

}
