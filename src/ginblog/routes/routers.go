package routes

import (
	"fmt"
	setting "ginblog/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() {
	// 选择模式，debug为开发模式，test为测试模式，release为生产模式
	gin.SetMode(setting.AppMode)

	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"go": "hello",
			})
		})
	}
	err := r.Run(setting.HttpPort)
	if err != nil {
		fmt.Println("gin服务启动失败", err.Error())
	}
}
