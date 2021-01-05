package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type test struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Sex  string `json:"sex"`
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

	r, u := selectOne(2)
	if r.Error != nil {
		log.Fatalln(r.Error)
	}
	fmt.Printf("共获取%v条数据：%v\n", r.RowsAffected, u)
	data, err := json.Marshal(u)
	if err != nil {
		fmt.Println("JSON序列化失败", err.Error())
	}
	fmt.Printf("JSON序列化成功：%s\n", data)
	r1, res := selectMax("男")
	if r1.Error != nil {
		log.Fatalln(r1.Error)
	}
	fmt.Printf("共获取%v条数据：%v\n", r1.RowsAffected, res)

	r2, res2 := selectIn("kyo", "JayChou", "JJLin")
	if r2.Error != nil {
		fmt.Println(r2.Error)
	}
	fmt.Printf("共获取%v条数据：%v\n", r2.RowsAffected, res2)

}

// 根据索引获取第一条匹配的数据
func selectOne(index int) (result *gorm.DB, user test) {
	user = test{}
	result = gOrmDB.Where("Id = ?", index).First(&user)
	return
}

// 根据索引获取所有匹配的数据
func selectMax(sex string) (result *gorm.DB, t []test) {
	user := test{}
	result = gOrmDB.Where("Sex = ?", sex).Find(&user)
	t = append(t, user)
	return
}

// 根据范围获取所有匹配的数据
func selectIn(v1, v2, v3 string) (result *gorm.DB, user test) {
	user = test{}
	result = gOrmDB.Where("name IN ?", []string{v1, v2, v3}).Find(&user)
	return
}
