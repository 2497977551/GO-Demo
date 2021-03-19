package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	setting "ginblog/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() {
	// 选择模式，debug为开发模式，test为测试模式，release为生产模式
	gin.SetMode(setting.AppMode)

	r := gin.New()
	r.Use(gin.Recovery(), middleware.Log(), middleware.Cors())

	rv1 := r.Group("Api/V1")
	rv1.Use(middleware.JWTTokenMid())
	{
		// 校验用户是否存在
		rv1.POST("UserNameCheck", v1.QueryUserIfExist)
		// 编辑用户
		rv1.POST("UpdateUser", v1.EditUser)
		// 删除用户
		rv1.POST("DeleteUser", v1.DeleteUser)
		// 添加文章
		rv1.POST("AddArticle", v1.AddOneArticle)
		//	编辑文章
		rv1.POST("UpdateArticle", v1.UpdateArticle)
		//	删除文章
		rv1.POST("RemoveArticle", v1.RemoveArticle)
		// 添加分类
		rv1.POST("AddCategory", v1.AppendCategory)
		// 修改分类
		rv1.POST("UpdateCategory", v1.UpdateCate)
		// 删除分类
		rv1.POST("DeleteCategory", v1.RemoveCate)
		//	上传文件
		rv1.POST("UploadFile", v1.Upload)
	}

	noMid := r.Group("Api/V1")
	{
		// 查询单个分类下所有文章
		noMid.GET("QueryOneCate", v1.QueryAllCateArticle)
		// 查询所有分类
		noMid.GET("QueryAllCategory", v1.QueryCategoryList)
		//	模糊查询文章
		noMid.GET("QueryArticle", v1.QueryOneArticle)
		//	查询文章列表
		noMid.GET("QueryArticleList", v1.QueryArticleList)
		// 查询单个用户
		noMid.GET("QueryUser", v1.QueryUser)
		// 查询所有用户
		noMid.GET("QueryAllUser", v1.QueryAllUserList)
		// 添加用户
		noMid.POST("AddUser", v1.AddUser)
		//	登录
		noMid.POST("Login", v1.UserLogin)
	}
	err := r.Run(setting.HttpPort)
	if err != nil {
		log.Fatalln("gin服务启动失败", err.Error())
	}
}
