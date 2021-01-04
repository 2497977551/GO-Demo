package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	type msg struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	r.GET("/", func(c *gin.Context) {
		data := msg{
			1,
			"Athena",
			16,
		}
		c.JSON(http.StatusOK, data)
	})
	r.Run(":1000")
}
