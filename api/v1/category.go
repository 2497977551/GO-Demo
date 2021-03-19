package v1

import (
	"ginblog/model"
	"ginblog/utils/ErrorInfo"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// 添加分类
func AppendCategory(c *gin.Context) {
	from := model.Category{}
	err = c.ShouldBindJSON(&from)
	if err != nil {
		log.Fatalln(err)
	} else {
		code := model.AddCategory(from)
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": ErrorInfo.GetErrMsg(code),
		})
	}

}

// 查询分类是否存在

// 查询单个分类下的所有文章
func QueryAllCateArticle(c *gin.Context) {
	id := c.Query("ID")
	data, code := model.QueryCateBlog(id)

	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"Message": ErrorInfo.GetErrMsg(code),
		"Data":    data,
	})
}

// 查询分类列表
func QueryCategoryList(c *gin.Context) {
	data, code := model.CategoryList()
	c.JSON(http.StatusOK, gin.H{
		"Status":  code,
		"Message": ErrorInfo.GetErrMsg(code),
		"Data":    data,
	})
}

// 编辑分类
func UpdateCate(c *gin.Context) {
	from := model.Category{}
	err = c.ShouldBindJSON(&from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status":  404,
			"Message": "参数错误",
		})
		log.Fatalln(err)
	} else {
		code := model.EditCate(from)
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": ErrorInfo.GetErrMsg(code),
		})
	}

}

// 删除分类
func RemoveCate(c *gin.Context) {
	from := model.SaveList{}
	err = c.ShouldBindJSON(&from)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"Status":  404,
			"Message": "参数错误",
		})
		log.Fatalln(err)
	} else {
		code := model.DeleteCate(from.Id)
		c.JSON(http.StatusOK, gin.H{
			"Status":  code,
			"Message": ErrorInfo.GetErrMsg(code),
		})
	}
}
