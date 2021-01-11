package main

// 关注用户
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type follows struct {
	MemberId int `json:"member_id"`
	UserId   int `json:"user_id"`
	Status   int `json:"status"`
}

func follow(memberId, userId int) (res *gorm.DB) {
	user := Follow{MemberId: memberId, UserId: userId}
	res = Db.Debug().Select("MemberId", "UserId").Create(&user)
	return
}
func removeFollow(memberId, userId int) (res *gorm.DB) {

	res = Db.Debug().Where("Member_Id = ? AND User_Id = ?", memberId, userId).Delete(&Follow{})
	return
}
func FollowMember(c *gin.Context) {
	from := follows{}
	if err := c.ShouldBindJSON(&from); err != nil {
		fmt.Println(err.Error())
		Data["code"] = 401
		Data["message"] = "类型错误！"
		c.JSON(http.StatusUnauthorized, gin.H{
			"data": Data,
		})
	} else {
		if from.Status == 0 {
			if res := follow(from.MemberId, from.UserId); res.Error != nil {
				fmt.Println(res.Error)
			} else {
				if res.RowsAffected == 1 {
					Data["code"] = 200
					Data["message"] = "关注成功！"
					c.JSON(http.StatusUnauthorized, gin.H{
						"data": Data,
					})
				} else {
					Data["code"] = 200
					Data["message"] = "关注失败！"
					c.JSON(http.StatusUnauthorized, gin.H{
						"data": Data,
					})
				}
			}
		} else if from.Status == 1 {
			if res := removeFollow(from.MemberId, from.UserId); res.Error != nil {
				fmt.Println(res.Error)
			} else {
				if res.RowsAffected == 1 {
					Data["code"] = 200
					Data["message"] = "取消关注成功！"
					c.JSON(http.StatusUnauthorized, gin.H{
						"data": Data,
					})
				} else {
					Data["code"] = 200
					Data["message"] = "取消关注失败！"
					c.JSON(http.StatusUnauthorized, gin.H{
						"data": Data,
					})
				}
			}
		} else {
			Data["code"] = 401
			Data["message"] = "错误，请检查参数是否填写正确！"
			c.JSON(http.StatusUnauthorized, gin.H{
				"data": Data,
			})
		}

	}
}
