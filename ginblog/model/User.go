package model

import (
	"encoding/base64"
	"ginblog/utils/ErrorInfo"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	Model
	UserName string `gorm:"column:UserName;NOT NULL" json:"UserName" validate:"required,min=4,max=10" label:"用户名"` // 用户名
	PassWord string `gorm:"column:PassWord;NOT NULL" json:"PassWord" validate:"required,min=8,max=16" label:"密码"`  // 用户密码
	Role     bool   `gorm:"column:Role;NOT NULL" json:"Role" validate:"required" label:"权限"`                       // 用户权限

}
type queryUser struct {
	Id       uuid.UUID `gorm:"column:ID;primary_key" json:"Id" validate:"required"` // 用户id
	UserName string    `gorm:"column:UserName;NOT NULL" json:"UserName" validate:"required"`
	// 用户名
	Role         bool      `gorm:"column:Role;NOT NULL" json:"Role" `       // 用户权限
	CreationTime time.Time `gorm:"column:CreationTime" json:"CreationTime"` // 创建时间
}

// 判断用户名是否存在
func CheckUser(username string) int {
	var users User

	db.Select("UserName").Where("UserName = ?", username).First(&users)

	if users.UserName != "" {

		return ErrorInfo.ERRUserNameExists
	}
	return ErrorInfo.SucCse
}

// 创建用户
func CreateUser(users *User) int {

	err = db.Omit("UpdateTime").Create(&users).Error
	if err != nil {
		return ErrorInfo.Error
	}
	if users.UserName == "" || users.PassWord == "" {
		return ErrorInfo.Error
	}
	users.PassWord = HashPwd(users.PassWord)
	return ErrorInfo.SucCse
}

// 查询用户列表，每次查询二十
func GetAllUser(page, pageSize int) (users []queryUser, id int64, p1, p2 int, code int) {

	if page > 0 && pageSize > 0 {

		db.Table("user").Where("DeleteTime is null").Offset((page - 1) * pageSize).Limit(pageSize).Find(&users)

		db.Debug().Table("user").Where("DeleteTime is null").Count(&id)
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
		err = db.Table("user").Where("UserName = ?", username).First(&users).Error
		if err != nil {
			code = ErrorInfo.ERRUserNoExistent
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

// 密码加密
// 密码加密钩子函数
func (u *User) BeforeSave() {
	u.PassWord = HashPwd(u.PassWord)
}
func HashPwd(password string) (pwd string) {
	salt := []byte{0xc8, 0x28, 0xf2, 0x58, 0xa7, 0x6a, 0xad, 0x7b}
	const KeyLin = 20
	dk, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, KeyLin)
	if err != nil {
		log.Fatal(err)
	}
	pwd = base64.StdEncoding.EncodeToString(dk)
	return
}

// 删除用户
func DeleteUsers(id uuid.UUID) int {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return ErrorInfo.Error
	}
	return ErrorInfo.SucCse
}

// 编辑用户
type Users struct {
	Id         uuid.UUID `gorm:"column:ID;primary_key" json:"Id" binding:"required"`          // 用户id
	UserName   string    `gorm:"column:UserName;NOT NULL" json:"UserName" binding:"required"` // 用户名
	Role       bool      `gorm:"column:Role;NOT NULL" json:"Role" binding:"required"`         // 用户权限
	UpdateTime time.Time `gorm:"column:UpdateTime" json:"UpdateTime"`                         // 修改时间
}

func beforeUpdate(id uuid.UUID) (g *gorm.DB) {
	var user User
	g = db.Unscoped().Table("user").Where("id = ? AND DeleteTime is not null", id).Find(&user)
	return
}
func UpdateUser(id uuid.UUID, u Users) int {
	var err error
	g := beforeUpdate(id)
	CheckUser(u.UserName)
	if g.RowsAffected != 0 {
		return ErrorInfo.Error
	}
	err = db.Table("user").Where("id = ?", id).Updates(Users{UserName: u.UserName, Role: u.Role, UpdateTime: time.Now()}).Error
	if err != nil {
		return ErrorInfo.Error
	}
	return ErrorInfo.SucCse
}

// 验证登录
func Login(name, pwd string) int {
	var user User
	db.Table("User").Where("UserName = ?", name).First(&user)
	if user.Id == uuid.Nil {
		return ErrorInfo.ERRUserNoExistent
	}
	if str := HashPwd(pwd); str != user.PassWord {

		return ErrorInfo.ERRPassWordWrong
	}
	if !user.Role {
		return ErrorInfo.ERRNoPermission
	}
	return ErrorInfo.SucCse
}
