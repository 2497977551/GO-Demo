package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id         uuid.UUID       `gorm:"column:ID;NOT NULL" json:"ID"`            // id
	CreatedAt  time.Time       `gorm:"column:CreationTime" json:"CreationTime"` // 创建时间
	UpdatedAt  time.Time       `gorm:"column:UpdateTime" json:"UpdateTime"`     // 更新时间
	DeleteTime *gorm.DeletedAt `gorm:"column:DeleteTime" json:"DeleteTime,omitempty"`
}

type SaveList struct {
	Id   uuid.UUID `gorm:"column:ID;NOT NULL"json:"id" `
	Name string    `gorm:"column:Name;NOT NULL"json:"name" `
}
type ArticleList struct {
	Model

	Title    string
	Describe string
	Content  string

	CID uuid.UUID
}
type EditArticles struct {
	ID        uuid.UUID `gorm:"column:ID" json:"ID"`
	Title     string    `gorm:"column:Title" json:"Title"`
	Describe  string    `gorm:"column:Describe" json:"Describe"`
	Content   string    `gorm:"column:Content" json:"Content"`
	CID       uuid.UUID `gorm:"column:CID" json:"CID"`
	UpdatedAt time.Time `gorm:"column:UpdateTime" json:"UpdateTime"`
}
type UserLogin struct {
	UserName string `gorm:"column:UserName;NOT NULL" json:"UserName" binding:"required"` // 用户名
	PassWord string `gorm:"column:PassWord;NOT NULL" json:"PassWord" binding:"required"` // 用户密码
}
