package main

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

func IndexDefer(g *gorm.DB) {
	sqlDb, err := g.DB()
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = sqlDb.Close()
	if err != nil {
		log.Fatalln("数据库连接断开失败", err.Error())
	} else {
		fmt.Println("数据库连接断开成功")
	}

}
