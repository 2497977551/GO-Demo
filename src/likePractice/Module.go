package main

import (
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB
var Data = make(map[string]interface{})

type Member struct {
	Id           uint      `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Username     string    `gorm:"column:username;NOT NULL" json:"username"`
	Password     string    `gorm:"column:password;NOT NULL" json:"password"`
	CreationTime time.Time `gorm:"column:creation_time;NOT NULL" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time;NOT NULL" json:"update_time"`
}
type Article struct {
	Id           uint      `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	MemberId     uint      `gorm:"column:member_id;AUTO_INCREMENT" json:"member_id"`
	Title        string    `gorm:"column:title" json:"title"`
	Content      string    `gorm:"column:content" json:"content"`
	PraiseNumber uint      `gorm:"column:praise_number" json:"praise_number"`
	CreationTime time.Time `gorm:"column:creation_time" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	DeleteTime   time.Time `gorm:"column:delete_time" json:"delete_time"`
}

type Notice struct {
	Id           uint      `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	MemberId     uint      `gorm:"column:member_id" json:"member_id"`
	Title        string    `gorm:"column:title" json:"title"`
	Content      string    `gorm:"column:content" json:"content"`
	CreationTime time.Time `gorm:"column:creation_time" json:"creation_time"`
	UpdateTime   time.Time `gorm:"column:update_time" json:"update_time"`
	DeleteDate   time.Time `gorm:"column:delete_date" json:"delete_date"`
}
