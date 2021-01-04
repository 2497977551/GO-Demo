package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
	中间件--请求过程中的处理函数
	比如身份认证
*/

// 定义中间件
func m1(c *gin.Context) {
	start := time.Now()
	c.Set("name", "Athena") // 设置请求上下文
	c.Next()                // 调用后续程序
	//c.Abort() 阻止调用后续程序
	end := time.Since(start)
	fmt.Printf("耗时：%v\n", end)
}

// 身份认证中间件(闭包)
//func authentication(onOff bool)gin.HandlerFunc{
// 连接数据库或其他操作
//	return func(c *gin.Context) {
//		// onOff是一个开关，代表是否校验
//		if onOff {
//			if 是登录状态 {
//				c.Next()
//			}else {
//				c.Abort()
//			}
//		}else {
//			c.Next()
//		}
//
//	}
//}
func main() {
	r := gin.Default()

	// 全局中间件,在所有请求前添加
	//r.Use(m1)

	// 单个请求中间件，中间件函数放到接口处理函数前即可
	r.GET("/index", m1, func(c *gin.Context) {
		// 获取请求上下文中的值，这么做可以跨路由取值
		name, ok := c.Get("name")
		if !ok {
			name = "用户不存在"
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "ok",
			"name": name,
		})
	})

	// 局部中间件 -- 给路由组添加中间件
	// 方法一
	userGroup := r.Group("/user", m1)
	// 方法二
	//userGroup.Use(m1)
	{
		userGroup.GET("/name", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name": "Joshua",
			})
		})
		userGroup.GET("/age", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"age": 18,
			})
		})
	}
	r.Run(":9090")
}
