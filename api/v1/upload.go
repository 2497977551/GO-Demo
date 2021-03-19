package v1

import (
	"ginblog/model"
	"ginblog/utils/ErrorInfo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrorInfo.Error,
			"Message": err.Error(),
		})
		return
	}

	url, code := model.UploadFile(file, fileHeader.Size)
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"Message": ErrorInfo.GetErrMsg(code),
		"Url":     url,
	})
	return
}
