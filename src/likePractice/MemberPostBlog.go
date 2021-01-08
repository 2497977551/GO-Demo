package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func posts(memberId uint, title, blog string, creationTime time.Time) (result *gorm.DB) {
	user := Article{MemberId: memberId, Title: title, Content: blog, CreationTime: creationTime}
	result = Db.Debug().Omit("Id", "PraiseNumber", "UpdateTime", "DeleteTime").Create(&user)
	return
}

func PostBlog(c *gin.Context) {
	form := Article{}
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println(err.Error())
		Data["code"] = 401
		Data["message"] = "类型错误！"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": Data,
		})
	} else {
		fmt.Println(form)
		res := posts(form.MemberId, form.Title, form.Content, time.Now())
		if res.Error != nil {
			fmt.Println(res.Error)
			Data["code"] = 401
			Data["message"] = "发布失败，请检查用户ID是否正确！"
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": Data,
			})
		} else {
			if res.RowsAffected == 1 {
				Data["code"] = 200
				Data["message"] = "发布成功！"
				c.JSON(http.StatusOK, gin.H{
					"data": Data,
				})
			} else {
				Data["code"] = 401
				Data["message"] = "发布失败，请检查是否输入正确！"
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": Data,
				})
			}
		}
	}
}
