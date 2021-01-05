package main

import (
	"fmt"
	"gorm.io/gorm"
)

var gOrmDb *gorm.DB

type test struct {
	Id   int    `gorm:"type:int(100);column:id;primary_key;AUTO_INCREMENT;uniqueIndex" json:"id"` // id
	Name string `gorm:"type:varchar(10)column:name;NOT NULL" json:"name"`                         // 姓名
	Sex  string `gorm:"type:char(1);column:sex;NOT NULL" json:"sex"`                              // 性别
}

func init() {
	var err error
	err, gOrmDb = IndexInit()
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	} else {
		fmt.Println("数据库连接成功")
	}
}
func main() {
	// 关闭数据库连接
	defer IndexDefer(gOrmDb)

	deleteOne(1, "joshua", "男")
}

// 删除单条记录
func deleteOne(id int, name, sex string) {
	user := test{Id: id, Name: name, Sex: sex}
	gOrmDb.Delete(&user)
}
