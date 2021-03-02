package v1

import (
	"ginblog/model"
	"ginblog/utils/ErrorInfo"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 添加文章
func AddOneArticle(c *gin.Context) {
	from := model.Article{}
	err := c.ShouldBindJSON(&from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Code":    404,
			"Message": "参数错误",
		})
		return
	}
	code := model.AddArticle(from.Title, from.Describe, from.Content, from.CID)
	c.JSON(http.StatusOK, gin.H{
		"Code":    code,
		"Message": ErrorInfo.GetErrMsg(code),
	})
}

// 查询单个文章
func QueryOneArticle(c *gin.Context) {
	title := c.Query("Title")
	data, code := model.QueryArticle(title)
	c.JSON(http.StatusOK, gin.H{
		"Data":    data,
		"Status":  code,
		"Message": ErrorInfo.GetErrMsg(code),
	})
}

// 查询文章列表
func QueryArticleList(c *gin.Context) {
	var (
		pageSize int
		pageNum  int
		perr     error
	)
	pageSize, perr = strconv.Atoi(c.PostForm("PageSize"))
	pageNum, perr = strconv.Atoi(c.PostForm("PageNum"))
	if perr != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status":  ErrorInfo.Error,
			"Message": "参数错误",
		})
		return
	}
	data, code, total := model.QueryAllArticle(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"Data":    data,
		"Total":   total,
		"Status":  code,
		"Message": ErrorInfo.GetErrMsg(code),
	})
	return
}

// 编辑文章
func UpdateArticle(c *gin.Context) {
	from := model.EditArticles{}
	err := c.ShouldBindJSON(&from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status":  404,
			"Message": err,
		})
	}
	code := model.EditArticle(from)
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"Message": ErrorInfo.GetErrMsg(code),
	})
}

// 删除文章
func RemoveArticle(c *gin.Context) {
	from := model.Article{}
	err := c.ShouldBindJSON(&from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status":  404,
			"Message": err.Error(),
		})
		return
	}
	code := model.DeleteArticle(from.Id)
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"Message": ErrorInfo.GetErrMsg(code),
	})
}
