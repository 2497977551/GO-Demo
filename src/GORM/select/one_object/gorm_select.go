package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type test struct {
	Id   int
	Name string
	Sex  string
}

var gOrmDB *gorm.DB

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name
	var err error
	dsn := "root:admin@tcp(localhost:3306)/GoWeb?charset=utf8mb4&parseTime=True&loc=Local"
	gOrmDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("数据库连接失败：", err.Error())
	} else {
		fmt.Println("数据库连接成功")
	}

}
func main() {
	defer func() {
		sqlDb, err := gOrmDB.DB()
		if err != nil {
			log.Fatalln(err.Error())
		}
		err = sqlDb.Close()
		if err != nil {
			log.Fatalln("数据库连接断开失败", err.Error())
		} else {
			fmt.Println("数据库连接断开成功")
		}
	}()

	var result *gorm.DB
	user := test{}

	// 根据主键查询第一条数据: SELECT * FROM tests ORDER BY id LIMIT 1;
	result = gOrmDB.First(&user)
	// 返回查询到的数量
	fmt.Println(result.RowsAffected)
	// 返回错误
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println("根据主键查询第一条数据:", user)

	user2 := test{}
	// 随机查询第一条数据
	result = gOrmDB.Take(&user2)
	// 返回查询到的数量
	fmt.Println(result.RowsAffected)
	fmt.Println("随机查询第一条数据:", user2)

	user3 := test{}
	// 根据主键查询最后一条数据
	result = gOrmDB.Last(&user3)
	// 返回查询到的数量
	fmt.Println(result.RowsAffected)
	fmt.Println("根据主键查询最后一条数据:", user3)

	user4 := test{}
	// 查询所有数据
	result = gOrmDB.Find(&user4)
	// 返回查询到的数量
	fmt.Println(result.RowsAffected)
	fmt.Println("查询所有数据:", user4)

}
