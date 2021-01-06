package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type userInfo struct {
	Id       int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Username string `gorm:"column:username;NOT NULL" json:"username"`
	Password string `gorm:"column:password;NOT NULL" json:"password"`
}

var gOrmDb *gorm.DB

func init() {
	var err error
	dsn := "root:admin@tcp(localhost:3306)/GoWeb?charset=utf8mb4&parseTime=True&loc=Local"
	gOrmDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("数据库连接失败：", err.Error())
	} else {
		fmt.Println("数据库连接成功")
	}
}
func main() {
	// 关闭数据库连接
	defer func() {
		sqlDb, err := gOrmDb.DB()
		if err != nil {
			log.Fatalln(err.Error())
		}
		err = sqlDb.Close()
		if err != nil {
			log.Fatalln("数据库连接断开失败", err.Error())
		} else {
			fmt.Println("数据库连接断开成功")
		}
	}()

	r := gin.Default()

	r.Use(Cors1(), cookieChecks()) // 设置中间件
	r.POST("/login", PostParm)     // POST请求
	_ = r.Run(":1000")             // 端口
}

// 返回函数
func PostParm(c *gin.Context) {

	form := &userInfo{}

	if c.BindJSON(&form) == nil {
		fmt.Println(form.Username, form.Password)
		r, u := selectOneTest(form.Username, form.Password)

		if r.Error != nil {
			fmt.Println(r.Error)
		}
		// 判断传入的参数在数据库是否存在
		if form.Username == u.Username && form.Password == u.Password {
			c.JSON(200, gin.H{
				"status":   "you are logged in",
				"UserName": u.Username,
				"Password": u.Password,
			})

		} else {
			c.JSON(401, gin.H{"status": "unauthorized"})
		}
	}
}
func selectOneTest(users, pwd string) (result *gorm.DB, user userInfo) {
	user = userInfo{}
	result = gOrmDb.Where("username = ? AND password = ?", users, pwd).First(&user)
	return
}

// 设置cookie
func cookieChecks() gin.HandlerFunc {
	return func(c *gin.Context) {

		if cookie, err := c.Request.Cookie("cookie"); err == nil {
			value := cookie.Value
			if _, err := strconv.Atoi(value); err == nil {

				s := RandChar(20)
				cookie.Value = fmt.Sprintf("%s", s)
			}
			http.SetCookie(c.Writer, cookie)
			c.Next()
		} else {
			cookie := &http.Cookie{
				Name:  "cookie",
				Value: "0",
			}
			http.SetCookie(c.Writer, cookie)
		}
	}
}

// 随机字符串cookie
const char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandChar(size int) string {
	rand.NewSource(time.Now().UnixNano()) // 产生随机种子
	var s bytes.Buffer
	for i := 0; i < size; i++ {
		s.WriteByte(char[rand.Int63()%int64(len(char))])
	}
	return s.String()
}

// 解决请求跨域
func Cors1() gin.HandlerFunc {
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
