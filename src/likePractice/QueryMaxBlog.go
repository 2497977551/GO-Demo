package main

// 查询所有博客
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type member struct {
	Id       uint   `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Username string `gorm:"column:username;NOT NULL" json:"username"`
}

func queryBlog(page, pageSize int) (res, res2 *gorm.DB, user []member, id int64, p1, p2 int) {

	if page > 0 && pageSize > 0 {
		res = Db.Debug().Table("member").Offset((page - 1) * pageSize).Limit(pageSize).Find(&user)
		fmt.Println("go:", time.Now())

		res2 = Db.Debug().Table("member").Count(&id)
		fmt.Println("no go", time.Now())
		p1, p2 = page, pageSize

	}
	return
}

func QueryMaxBlog(c *gin.Context) {
	from := Paging{}
	if err := c.ShouldBindJSON(&from); err != nil {
		fmt.Println(err.Error())
		Data["code"] = 401
		Data["message"] = "类型错误！"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": Data,
		})
	} else {

		if res, res2, user, count, page, pageSize := queryBlog(from.Page, from.PageSize); res.Error != nil && res2.Error != nil {
			fmt.Println(res.Error, res2.Error)
			Data["code"] = 401
			Data["message"] = "查询失败，请检查是否填写正确！"
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": Data,
			})
		} else {
			if res.RowsAffected > 0 {
				Data["code"] = 200
				Data["message"] = "查询成功！"
				Data["items"] = user
				Data["count"] = count
				Data["currentPage"] = page
				Data["pageNumber"] = pageSize
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": Data,
				})

			} else {
				data := make(map[string]interface{})
				data["code"] = 401
				data["message"] = "查询失败，请检查参数是否填写正确！"
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": data,
				})

			}
		}
	}
}
