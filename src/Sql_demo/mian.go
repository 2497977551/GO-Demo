package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func main() {

	db, _ = sql.Open("mysql", "root:admin@tcp(localhost:3306)/GoWeb")
	//设置数据库最大连接数
	db.SetConnMaxLifetime(10)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(5)
	err := db.Ping()
	if err != nil {
		fmt.Println("MySql数据库连接失败", err.Error())
	} else {
		fmt.Println("MySql连接成功")
	}

	one, e := getOne("Athena")
	if e != nil {
		fmt.Println("查询错误", e.Error())
	}
	fmt.Println(one)
	one.name += "123"

	v, e2 := getWhole()
	if e2 != nil {
		log.Fatalln(e2.Error())
	}
	fmt.Println(v)

	e3 := one.update()
	if e3 != nil {
		fmt.Println("查询错误", e.Error())
	}
	one1, _ := getOne("Athena")
	fmt.Println(one1)
}
