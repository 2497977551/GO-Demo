package model

import uuid "github.com/satori/go.uuid"

type SaveList struct {
	Id   uuid.UUID `gorm:"column:ID;NOT NULL" `
	Name string    `gorm:"column:Name;NOT NULL" `
}
