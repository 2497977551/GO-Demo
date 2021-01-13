package model

import "time"

type User struct {
	Id           string     `gorm:"column:ID;primary_key" json:"Id"`          // 用户id
	UserName     string     `gorm:"column:UserName;NOT NULL" json:"UserName"` // 用户名
	PassWord     string     `gorm:"column:PassWord;NOT NULL" json:"PassWord"` // 用户密码
	Role         int        `gorm:"column:Role;NOT NULL" json:"Role"`         // 用户权限
	CreationTime *time.Time `gorm:"column:CreationTime" json:"CreationTime"`  // 创建时间
	UpdateTime   *time.Time `gorm:"column:UpdateTime" json:"UpdateTime"`      // 修改时间
}
