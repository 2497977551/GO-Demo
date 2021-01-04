package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
	1.路由分组
	1.拥有共同URL前缀的路由划分为一个路由组。
	2.一对{}包裹同组的路由。

	2.NoRouter
	即不存在的路由，当用户输入不存在的路由时的处理
*/
func main() {
	r := gin.Default()

	// 路由组就是将拥有相同的父级路由分组,如：/user/name或/user/age
	userGroup := r.Group("/User")
	{
		userGroup.GET("/Name", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name": "joshua",
			})
		})
		userGroup.GET("/Age", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"age": 18,
			})
		})
	}

	// 路由组嵌套
	shopGroup := r.Group("/Furniture")
	{
		shopGroup.GET("/Table", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name": "桌子",
			})
		})
		shopGroup.GET("/Chair", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"name": "椅子",
			})
		})
		// 路由组中嵌套路由组
		aea := shopGroup.Group("/AnElectricAppliance")
		{
			aea.GET("/ElectricFan", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"name": "电风扇",
				})
			})
		}
	}
	r.LoadHTMLFiles("./NoRouter.html")
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "NoRouter.html", nil)
	})

	r.Run(":9090")
}
