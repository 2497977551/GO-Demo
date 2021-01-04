package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/web", func(c *gin.Context) {
		//name := c.Query("query") // 通过key取值
		// 当找不到query这个key的时候就取第二个参数的默认值
		name := c.DefaultQuery("query", "默认值")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})
	r.Run(":1000")
}
