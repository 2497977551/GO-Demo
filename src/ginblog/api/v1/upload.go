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

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrorInfo.Error,
			"Message": err.Error(),
		})
		return
	}

	if file.Size >= 1024*1024*3 {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrorInfo.Error,
			"Message": "文件已大于3MB",
		})
		fmt.Println(file.Size)
		return
	}
	fileName := file.Filename
	url, code := model.UploadFile(fileName)
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"Message": ErrorInfo.GetErrMsg(code),
		"Url":     url,
	})
	return
}
