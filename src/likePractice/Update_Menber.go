package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func update(id uint, password string, updateTime time.Time) (res *gorm.DB) {
	user := Member{}
	res = Db.Debug().Model(&user).Select("Password", "UpdateTime").Where("id = ?", id).Updates(map[string]interface{}{
		"Password":   password,
		"UpdateTime": updateTime,
	})
	return
}

func UpdateMembers(c *gin.Context) {
	form := Member{}
	data := make(map[string]interface{})
	if err := c.BindJSON(&form); err != nil {
		fmt.Println(err.Error())

	}

	if res := update(form.Id, form.Password, time.Now()); res.Error != nil {
		fmt.Println(res.Error)

	} else if res.RowsAffected == 1 {
		data["code"] = 200
		data["message"] = "修改成功！"
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})

	} else {
		data["code"] = 401
		data["message"] = "修改失败，请检查是否填写正确！"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": data,
		})
	}

}
