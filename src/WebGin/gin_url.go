package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})

	})
	r.Run(":9090")
}
