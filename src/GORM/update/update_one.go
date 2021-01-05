package main

import (
	"fmt"
	"gorm.io/gorm"
)

type test struct {
	Id   uint `gorm:"primaryKey"`
	Name string
	Sex  string
}

var gOrmDB *gorm.DB

func init() {
	var err error
	err, gOrmDB = IndexInit()
	if err != nil {
		fmt.Println("数据库连接失败...", err.Error())
	} else {
		fmt.Println("数据库连接成功")
	}
}
func main() {
	// 释放连接
	defer IndexDefer(gOrmDB)

	updateOne("Sex", "男")

	indexUpdateOne(5, "Name", "LeeHom")

	r1 := indexUpdateMap(4, "AKL", "女")
	fmt.Printf("indexUpdateMap更新了%v条记录\n", r1)

	r2 := indexUpdateStruct(3, "King", "女")
	fmt.Printf("indexUpdateStruct更新了%v条记录\n", r2)
}

// 更新单个字段
func updateOne(k, v string) {
	user := test{Id: 1, Name: "joshua", Sex: "0"}
	gOrmDB.Model(&user).Update(k, v)
}

// 根据条件更新单个字段
func indexUpdateOne(i int, k, v string) {
	user := test{}
	gOrmDB.Model(&user).Where("Id = ?", i).Update(k, v)
}

// 通过条件与 `map` 更新多个字段，不会更新零值字段
func indexUpdateMap(i uint, v1, v2 interface{}) (updateLen *gorm.DB) {
	user := test{}
	gOrmDB.Model(&user).Where("Id = ?", i).Updates(map[string]interface{}{"Name": v1, "Sex": v2})
	return
}

// 通过条件与 `struct` 更新多个字段，不会更新零值字段
func indexUpdateStruct(i int, v1, v2 string) (updateLen *gorm.DB) {
	user := test{}
	updateLen = gOrmDB.Model(&user).Where("Id= ?", i).Updates(test{Name: v1, Sex: v2})
	return
}
