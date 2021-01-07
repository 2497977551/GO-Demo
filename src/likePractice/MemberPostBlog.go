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
	result = Db.Debug().Omit("PraiseNumber", "UpdateTime", "DeleteTime").Create(&user)
	return
}

func PostBlog(c *gin.Context) {
	form := Article{}
	data := make(map[string]interface{})
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println(err.Error())
	} else {
		res := posts(form.MemberId, form.Title, form.Content, time.Now())
		if res.Error != nil {
			fmt.Println(res.Error)
		} else {
			if res.RowsAffected == 1 {
				data["code"] = 200
				data["message"] = "发布成功！"
				c.JSON(http.StatusOK, gin.H{
					"data": data,
				})
			} else {
				data["code"] = 401
				data["message"] = "发布失败，请检查是否输入是否正确！"
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": data,
				})
			}
		}
	}
}
