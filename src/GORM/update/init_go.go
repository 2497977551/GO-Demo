package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func IndexInit() (err error, gOrmDB *gorm.DB) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:admin@tcp(localhost:3306)/GoWeb?charset=utf8mb4&parseTime=True&loc=Local"
	gOrmDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return
}
