package model

import "time"

type Category struct {
	Id           string     `gorm:"column:ID;NOT NULL" json:"ID"`            // id
	Name         string     `gorm:"column:Name;NOT NULL" json:"Name"`        // 文章类别名称
	CreationTime *time.Time `gorm:"column:CreationTime" json:"CreationTime"` // 创建时间
	UpdateTime   *time.Time `gorm:"column:UpdateTime" json:"UpdateTime"`     // 修改时间
}
