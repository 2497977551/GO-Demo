package main

// 评论博客
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func insertBlog(memberId, articleId uint, comment string, creationTime time.Time) (res *gorm.DB) {
	user := Release{MemberId: memberId, ArticleId: articleId, Comment: comment, CreationTime: creationTime}
	res = Db.Debug().Select("MemberId", "ArticleId", "Comment", "CreationTime").Create(&user)
	return
}

func CommentBlog(c *gin.Context) {
	from := Release{}
	if err := c.ShouldBindJSON(&from); err != nil {
		fmt.Println(err.Error())
		Data["code"] = 401
		Data["message"] = "类型错误！"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": Data,
		})
	} else {
		res := insertBlog(from.MemberId, from.ArticleId, from.Comment, time.Now())
		if res.Error != nil {
			fmt.Println(res.Error)
			Data["code"] = 401
			Data["message"] = "发布失败，请检查是否填写正确！"
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": Data,
			})
		} else {
			if res.RowsAffected == 1 {
				Data["code"] = 200
				Data["message"] = "发布成功！"
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": Data,
				})
			} else {
				Data["code"] = 401
				Data["message"] = "发布失败，请检查是否填写正确！"
				c.JSON(http.StatusUnauthorized, gin.H{
					"data": Data,
				})
			}
		}
	}
}
