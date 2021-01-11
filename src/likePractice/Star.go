package main

// 点赞
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type getStar struct {
	Id     int `json:"id"`
	Choice int `json:"choice"`
}

func userStar(is, id int) (res *gorm.DB, whether bool) {
	if is == 0 {
		res = Db.Debug().Table("article").Where("id = ?", id).Update("praise_number", gorm.Expr("praise_number - ?", 1))
		whether = true
	} else if is == 1 {
		type article struct {
			PraiseNumber int `gorm:"column:praise_number" json:"praise_number"`
		}
		user := article{}
		Db.Debug().Select("praise_number").Where("id = ?", id).Find(&user)
		if user.PraiseNumber != 0 {
			res = Db.Debug().Table("article").Where("id = ?", id).Update("praise_number", gorm.Expr("praise_number + ?", 1))
			whether = true
		} else if user.PraiseNumber == 0 {
			res = Db.Debug().Table("article").Where("id = ?", id).Update("praise_number", 1)
			whether = true
		} else {
			whether = false
		}

	} else {
		whether = false
	}
	return
}

func Star(c *gin.Context) {
	from := getStar{}
	if err := c.ShouldBindJSON(&from); err != nil {
		fmt.Println(err.Error())
		Data["code"] = 401
		Data["message"] = "请仔细检查！"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": Data,
		})
	} else {
		if res, is := userStar(from.Choice, from.Id); res.Error != nil {
			Data["code"] = 401
			Data["message"] = "失败，请检查是否正确！"
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": Data,
			})
		} else {
			if is == true && res.RowsAffected == 1 {
				Data["code"] = 200
				Data["message"] = "点赞成功！"
				c.JSON(http.StatusOK, gin.H{
					"data": Data,
				})
			} else if is == false && res.RowsAffected == 1 {
				Data["code"] = 200
				Data["message"] = "取消成功！"
				c.JSON(http.StatusOK, gin.H{
					"data": Data,
				})
			} else {
				Data["code"] = 401
				Data["message"] = "失败！"
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": Data,
				})
			}
		}
	}
}
