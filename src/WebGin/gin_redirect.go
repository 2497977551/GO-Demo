package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	请求重定向与转发
*/
func main() {
	r := gin.Default()

	// 重定向
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	//	转发
	r.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		r.HandleContext(c)
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b",
		})
	})
	r.Run(":9090")
}
