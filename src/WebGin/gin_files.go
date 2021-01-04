package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

/*
	文件上传
*/
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./upload.html", "./uploads.html")
	r.GET("/files", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", nil)
	})

	r.GET("/MaxFiles", func(c *gin.Context) {
		c.HTML(http.StatusOK, "uploads.html", nil)
	})

	/*
		单文件上传
	*/
	// 限制上传文件的最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		// 从请求中获取文件
		file, err := c.FormFile("files")
		if err != nil {
			c.String(500, "上传图片错误")
		} else {
			//	保存文件到本地
			filepath := path.Join("./img", file.Filename)
			err := c.SaveUploadedFile(file, filepath)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, gin.H{
					"message":  "ok",
					"fileName": file.Filename,
				})
			}

		}
	})

	/*
		多文件上传
	*/
	r.POST("/uploads", func(c *gin.Context) {
		// 从请求中获取所有文件
		file, err := c.MultipartForm()
		if err != nil {
			c.String(500, "上传图片错误")
		} else {
			//	获取文件保存到本地
			file := file.File["files"]
			for index, file := range file {
				filepath := fmt.Sprintf("./img/%s_%d", file.Filename, index)
				err := c.SaveUploadedFile(file, filepath)
				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"message": err.Error(),
					})
				}

			}

			c.JSON(http.StatusOK, gin.H{
				"message":    "ok",
				"fileNumber": len(file),
			})

		}

	})
	r.Run(":9090")
}
