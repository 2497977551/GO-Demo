package main

//查询所有关注
import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type getId struct {
	Id int `json:"id"`
}

func queryFollow(id int) (res *gorm.DB, count int64) {

	res = Db.Debug().Table("follow").Where("user_id = ?", id).Count(&count)
	return
}

func QueryAllFollow(c *gin.Context) {
	from := getId{}
	if err := c.ShouldBindJSON(&from); err != nil {
		fmt.Println(err.Error())
		Data["code"] = 401
		Data["message"] = "请仔细检查！"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": Data,
		})
	} else {
		if res, count := queryFollow(from.Id); res.Error != nil {
			fmt.Println(res.Error)
			Data["code"] = 401
			Data["message"] = "查询失败，请检查是否填写正确！"
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": Data,
			})
		} else {
			if count >= 0 {
				Data["code"] = 200
				Data["message"] = "查询成功！"
				Data["count"] = count
				c.JSON(http.StatusOK, gin.H{
					"data": Data,
				})
			} else {
				Data["code"] = 401
				Data["message"] = "查询失败，请检查是否填写正确！"
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": Data,
				})
			}
		}
	}
}
