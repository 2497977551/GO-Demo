package main

import (
	"fmt"
	"ginblog/model"
	"ginblog/routes"
)

func main() {
	sqlDb := model.InitDb()
	routes.InitRouter()
	err := sqlDb.Close()
	if err != nil {
		fmt.Println("MySQL数据库断开失败", err)
	}
}
