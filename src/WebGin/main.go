package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "1000",
		"data": "hello golang",
	})
}
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	r.GET("/go", hello)
	r.Run(":1000")
}
