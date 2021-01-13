package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// 查询用户是否存在
func QueryUserIfExist(c *gin.Context) {
	fmt.Println("查询用户是否存在")
}

// 添加用户
func AddUser(c *gin.Context) {}

// 查询单个用户
func QueryUser(c *gin.Context) {}

// 查询所有用户列表
func QueryAllUserList(c *gin.Context) {

}

// 编辑用户
func EditUser(c *gin.Context) {}

// 删除用户
func DeleteUser(c *gin.Context) {}
