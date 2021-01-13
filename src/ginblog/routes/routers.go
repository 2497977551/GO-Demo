package routes

import (
	"fmt"
	v1 "ginblog/api/v1"
	setting "ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 选择模式，debug为开发模式，test为测试模式，release为生产模式
	gin.SetMode(setting.AppMode)

	r := gin.Default()

	rv1 := r.Group("Api/V1")
	{
		// 用户模块路由组
		v1user := rv1.Group("/User")
		{
			// 校验用户是否存在
			v1user.POST("/UserNameCheck", v1.QueryUserIfExist)
			// 添加用户
			v1user.POST("/AddUser", v1.AddUser)
			// 查询单个用户
			v1user.GET("/QueryUser", v1.QueryUser)
			// 查询所有用户
			v1user.GET("/QueryAllUser", v1.QueryAllUserList)
			// 编辑用户
			v1user.POST("/UpdateUser", v1.EditUser)
			// 删除用户
			v1user.POST("/DeleteUser", v1.DeleteUser)
		}
	}
	err := r.Run(setting.HttpPort)
	if err != nil {
		fmt.Println("gin服务启动失败", err.Error())
	}
}
