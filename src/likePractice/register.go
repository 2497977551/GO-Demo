package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

type Member struct {
	Id           int       `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Username     string    `gorm:"column:username;NOT NULL" json:"username"`
	Password     string    `gorm:"column:password;NOT NULL" json:"password"`
	CreationTime time.Time `gorm:"column:creation_time;NOT NULL" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time;NOT NULL" json:"update_time"`
}

var gOrmDb *gorm.DB

func init() {
	var err error
	err, gOrmDb = IndexInit()
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	} else {
		fmt.Println("连接成功")
	}
}
func main() {
	defer IndexDefer(gOrmDb)

	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		u := Member{}
		data := make(map[string]interface{})

		if err := c.BindJSON(&u); err != nil {

			fmt.Println(err.Error())

		} else {

			if u.Username != "" && u.Password != "" {
				registerUser(u.Username, u.Password, time.Now(), time.Now())
				data["status"] = "OK!"
				c.JSON(http.StatusOK, gin.H{
					"data": data,
				})

			} else {
				data["status"] = "error"
				c.JSON(http.StatusNotFound, gin.H{
					"data": data,
				})
			}
		}

	})

	err := r.Run(":5050")

	if err != nil {
		fmt.Println("服务启动成功")
	}

}

// 创建数据函数
func registerUser(username, password string, creationTime, updateTime time.Time) {
	user := Member{
		Username:     username,
		Password:     password,
		CreationTime: creationTime,
		UpdateTime:   updateTime,
	}
	gOrmDb.Create(&user)
}
