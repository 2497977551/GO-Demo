package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	r := gin.Default()
	r.Use(cookieCheck())
	r.GET("/http_cookie", func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("cookie"); err == nil {
			fmt.Println("http cookie")
			c.String(http.StatusOK, cookie.Value)
		}
	})
	_ = r.Run(":2020")
}

func cookieCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("cookie"); err == nil {
			value := cookie.Value
			if v, err := strconv.Atoi(value); err == nil {
				i := v + 1
				cookie.Value = fmt.Sprintf("%d", i)
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
