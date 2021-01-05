package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type test struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Sex  string `json:"sex"`
}

var gOrmDB *gorm.DB

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
	var err error
	dsn := "root:admin@tcp(localhost:3306)/GoWeb?charset=utf8mb4&parseTime=True&loc=Local"
	gOrmDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("数据库连接失败：", err.Error())
	} else {
		fmt.Println("数据库连接成功")
	}

}
func main() {
	r := gin.Default()
	r.Use(Cors())
	r.GET("/test", func(c *gin.Context) {
		r, v := selectOne(2)
		if r.Error != nil {
			fmt.Println(r.Error)
		}
		c.JSON(http.StatusOK, gin.H{
			"id":   v.Id,
			"name": v.Name,
			"sex":  v.Sex,
		})

	})
	err := r.Run(":5050")
	if err != nil {
		fmt.Println(err.Error())
	}
}
func selectOne(index int) (result *gorm.DB, user test) {
	user = test{}
	result = gOrmDB.Where("Id = ?", index).First(&user)
	return
}

// 跨域解决中间件
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}
