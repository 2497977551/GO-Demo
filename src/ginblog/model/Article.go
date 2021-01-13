package model

import "time"

type Article struct {
	Id           string     `gorm:"column:ID;NOT NULL" json:"ID"`             // 文章id
	Title        string     `gorm:"column:Title;NOT NULL" json:"Title"`       // 文章标题
	CID          string     `gorm:"column:CID;NOT NULL" json:"CID"`           // 文章分类
	Describe     string     `gorm:"column:Describe;NOT NULL" json:"Describe"` // 文章详情
	Content      string     `gorm:"column:Content;NOT NULL" json:"Content"`   // 文章主体内容
	CreationTime *time.Time `gorm:"column:CreationTime" json:"CreationTime"`  // 创建时间
	UpdateTime   *time.Time `gorm:"column:UpdateTime" json:"UpdateTime"`      // 更新时间
}
