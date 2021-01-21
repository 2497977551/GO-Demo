package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Id           uuid.UUID       `gorm:"column:ID;NOT NULL" json:"ID"`            // id
	CreationTime time.Time       `gorm:"column:CreationTime" json:"CreationTime"` // 创建时间
	UpdateTime   *time.Time      `gorm:"column:UpdateTime" json:"UpdateTime"`     // 修改时间
	DeleteTime   *gorm.DeletedAt `gorm:"column:DeleteTime" json:"DeleteTime"`
}
