package model

import (
	"database/sql"
	"fmt"
	"ginblog/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var (
	db  *gorm.DB
	err error
)

func InitDb() (sqlDb *sql.DB) {
	db, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//	//TablePrefix: "t_",   // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名
		},
	})
	if err != nil {
		fmt.Println("MySQL数据库连接失败", err.Error())
	}

	sqlDb, err = db.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。

	sqlDb.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDb.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDb.SetConnMaxLifetime(10 * time.Second)

	return
}
