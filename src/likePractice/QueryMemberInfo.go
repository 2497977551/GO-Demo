package main

// 查询用户信息

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func queryInfo(id uint) (res *gorm.DB, user Member) {
	user = Member{}
	res = Db.Debug().Where("id = ?", id).First(&user)
	return
}

func QueryMemberInfo(c *gin.Context) {
	from := Member{}
	if err := c.ShouldBindJSON(&from); err != nil {
		fmt.Println(err.Error())
		Data["code"] = 401
		Data["message"] = "类型错误！"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": Data,
		})
	} else {
		if res, user := queryInfo(from.Id); res.Error != nil {
			fmt.Println(res.Error)
			Data["code"] = 401
			Data["message"] = "发布失败，请检查是否填写正确！"
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": Data,
			})
		} else {
			userInfo := make(map[string]interface{})
			if res.RowsAffected == 1 {
				Data["code"] = 200
				Data["message"] = "查询成功！"
				userInfo["userName"] = user.Username
				Data["userInfo"] = userInfo
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": Data,
				})
			} else {
				Data["code"] = 401
				Data["message"] = "发布失败，请检查是否填写正确！"
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": Data,
				})
			}
		}
	}
}
