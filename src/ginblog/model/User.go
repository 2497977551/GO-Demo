package model

import (
	"ginblog/utils/ErrorInfo"
	uuid "github.com/satori/go.uuid"
	"time"
)

type User struct {
	Id           uuid.UUID `gorm:"column:ID;primary_key" json:"Id" binding:"required"`          // 用户id
	UserName     string    `gorm:"column:UserName;NOT NULL" json:"UserName" binding:"required"` // 用户名
	PassWord     string    `gorm:"column:PassWord;NOT NULL" json:"PassWord" binding:"required"` // 用户密码
	Role         bool      `gorm:"column:Role;NOT NULL" json:"Role" `                           // 用户权限
	CreationTime time.Time `gorm:"column:CreationTime" json:"CreationTime"`                     // 创建时间
	UpdateTime   time.Time `gorm:"column:UpdateTime" json:"UpdateTime"`                         // 修改时间
}
type queryUser struct {
	Id       uuid.UUID `gorm:"column:ID;primary_key" json:"Id" binding:"required"` // 用户id
	UserName string    `gorm:"column:UserName;NOT NULL" json:"UserName" binding:"required"`
	// 用户名
	Role         bool      `gorm:"column:Role;NOT NULL" json:"Role" `       // 用户权限
	CreationTime time.Time `gorm:"column:CreationTime" json:"CreationTime"` // 创建时间
}

// 判断用户名是否存在
func CheckUser(username string) int {
	var users User
	err = db.Debug().Select("UserName").Where("UserName = ?", username).First(&users).Error
	if users.UserName != " " && err != nil {
		return ErrorInfo.ERRUserNameExists
	}
	return ErrorInfo.SucCse
}

// 创建用户
func CreateUser(users *User) int {
	err = db.Debug().Omit("UpdateTime").Create(&users).Error
	if err != nil {
		return ErrorInfo.Error
	}
	return ErrorInfo.SucCse
}

// 查询用户列表，每次查询二十
func GetAllUser(page, pageSize int) (users []queryUser, id int64, p1, p2 int, code int) {

	if page > 0 && pageSize > 0 {
		db.Debug().Table("user").Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)

		db.Debug().Table("user").Count(&id)
		p1, p2 = page, pageSize
		code = ErrorInfo.SucCse
	} else {
		code = ErrorInfo.Error
	}
	return
}

// 查询单个用户
func QueryUsers(username string) (users []queryUser, code int) {
	if username != " " {
		err = db.Debug().Table("user").Where("UserName = ?", username).First(&users).Error
		if err != nil {
			code = ErrorInfo.Error
			return
		} else {
			code = ErrorInfo.SucCse
			return
		}

	} else {
		code = ErrorInfo.Error
		return
	}

}
