package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math"
	"os"
	"time"
)

func Log() gin.HandlerFunc {
	logger := logrus.New()
	return func(c *gin.Context) {
		starTime := time.Now()
		c.Next()
		endTime := time.Since(starTime)
		// 耗时
		takeTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(endTime.Nanoseconds()*100000.0))))
		// 主机名
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "未知"
		}
		// 状态码
		statusCode := c.Writer.Status()
		// ip
		clientIp := c.ClientIP()
		// 浏览器
		userAgent := c.Request.UserAgent()
		// 大小
		dataSize := c.Writer.Size()
		// 方法
		method := c.Request.Method
		// url
		path := c.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"Status":    statusCode,
			"Ip":        clientIp,
			"Agent":     userAgent,
			"Size":      dataSize,
			"Method":    method,
			"Path":      path,
			"SpendTime": takeTime,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode > 500 {
			entry.Error()
		} else if statusCode > 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
