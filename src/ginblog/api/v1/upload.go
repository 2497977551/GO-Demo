package v1

import (
	"fmt"
	"ginblog/model"
	"ginblog/utils/ErrorInfo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, err := c.FormFile("file")
	fileName := file.Filename
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrorInfo.Error,
			"Message": err.Error(),
		})
		return
	}
	if file.Size >= 1024*1024*5 {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrorInfo.Error,
			"Message": "文件过大",
		})
		fmt.Println(file.Size)
		return
	}
	url, code := model.UploadFile(fileName)
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"Message": ErrorInfo.GetErrMsg(code),
		"Url":     url,
	})
	return
}
