package main

// 发布广告
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func notice(memberId uint, title, content string, creationTime time.Time) (result *gorm.DB) {
	defer GetRecover()
	user := Notice{MemberId: memberId, Title: title, Content: content, CreationTime: creationTime}
	result = Db.Debug().Select("MemberId", "Title", "Content", "CreationTime").Create(&user)
	return
}

func MemberPostNotice(c *gin.Context) {
	defer GetRecover()
	from := Notice{}
	if err := c.ShouldBindJSON(&from); err != nil {
		fmt.Println(err.Error())
		Data["code"] = 401
		Data["message"] = "类型错误"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": Data,
		})
	} else {
		if res := notice(from.MemberId, from.Title, from.Content, time.Now()); res.Error != nil {
			fmt.Println(res.Error)
			Data["code"] = 401
			Data["message"] = "发布失败"
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": Data,
			})
		} else {
			if res.RowsAffected == 1 {
				Data["code"] = 200
				Data["message"] = "发布成功"
				c.JSON(http.StatusOK, gin.H{
					"data": Data,
				})
			} else {
				Data["code"] = 401
				Data["message"] = "发布失败"
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": Data,
				})
			}
		}
	}
}
