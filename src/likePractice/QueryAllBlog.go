package main

// 查询所有关注的人博客
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type article struct {
	Id           uint      `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	MemberId     uint      `gorm:"column:member_id;AUTO_INCREMENT" json:"member_id"`
	Title        string    `gorm:"column:title" json:"title"`
	Content      string    `gorm:"column:content" json:"content"`
	PraiseNumber uint      `gorm:"column:praise_number" json:"praise_number"`
	CreationTime time.Time `gorm:"column:creation_time" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	DeleteTime   time.Time `gorm:"column:delete_time" json:"-"`
}
type index struct {
	Id uint `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
}

func queryAll(id uint) (res, res2 *gorm.DB, article []article) {
	type follow struct {
		UserId int `gorm:"column:user_id" json:"user_id"`
	}
	var follows []follow
	res2 = Db.Debug().Where("member_id = ?", id).Find(&follows)
	for i, i2 := range follows {
		fmt.Println("key", i, follows[i])

		res = Db.Debug().Where("member_id = ?", i2.UserId).Order("creation_time DESC").Find(&article)

	}

	return
}
func QueryAllBlog(c *gin.Context) {
	from := index{}
	if err := c.ShouldBindJSON(&from); err != nil {
		fmt.Println(err.Error())
		Data["code"] = 401
		Data["message"] = "类型错误！"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": Data,
		})
	} else {
		if res, res2, user := queryAll(from.Id); res.Error != nil && res2.Error != nil {
			fmt.Println(res.Error)
			fmt.Println(res2.Error)
			Data["code"] = 401
			Data["message"] = "查询失败，请检查参数是否填写正确！"
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": Data,
			})
		} else {
			if res.RowsAffected != 0 && user != nil {
				Data["code"] = 200
				Data["message"] = "查询成功！"
				Data["blog"] = user
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": Data,
				})
			} else {
				Data["code"] = 400
				Data["message"] = "查询失败，请检查是否填写正确！"
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": Data,
				})
			}
		}
	}
}
