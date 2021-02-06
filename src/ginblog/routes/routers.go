package routes

import (
	"fmt"
	v1 "ginblog/api/v1"
	setting "ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	// 选择模式，debug为开发模式，test为测试模式，release为生产模式
	gin.SetMode(setting.AppMode)

	r := gin.Default()

	rv1 := r.Group("Api/V1")
	{
		// 用户模块路由组
		v1user := rv1.Group("User")
		{
			// 校验用户是否存在
			v1user.POST("UserNameCheck", v1.QueryUserIfExist)
			// 添加用户
			v1user.POST("AddUser", v1.AddUser)
			// 查询单个用户
			v1user.GET("QueryUser", v1.QueryUser)
			// 查询所有用户
			v1user.GET("QueryAllUser", v1.QueryAllUserList)
			// 编辑用户
			v1user.POST("UpdateUser", v1.EditUser)
			// 删除用户
			v1user.POST("DeleteUser", v1.DeleteUser)
		}

		// 文章模块路由组
		v1article := rv1.Group("Article")
		{
			// 添加文章
			v1article.POST("AddArticle", v1.AddOneArticle)
			//	模糊查询文章
			v1article.GET("QueryArticle", v1.QueryOneArticle)
			//	查询文章列表
			v1article.GET("QueryArticleList", v1.QueryArticleList)
			//	编辑文章
			v1article.POST("UpdateArticle", v1.UpdateArticle)
			//	删除文章
			v1article.POST("RemoveArticle", v1.RemoveArticle)
		}

		// 分类模块路由组
		v1category := rv1.Group("Category")
		{
			// 添加分类
			v1category.POST("AddCategory", v1.AppendCategory)
			// 查询所有分类
			v1category.GET("QueryAllCategory", v1.QueryCategoryList)
			// 修改分类
			v1category.POST("UpdateCategory", v1.UpdateCate)
			// 删除分类
			v1category.POST("DeleteCategory", v1.RemoveCate)
			// 查询单个分类下所有文章
			v1category.GET("QueryOneCate", v1.QueryAllCateArticle)
		}
	}
	err := r.Run(setting.HttpPort)
	if err != nil {
		fmt.Println("gin服务启动失败", err.Error())
	}
}
