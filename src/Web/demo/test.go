package main

//
//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//	"log"
//)
//
//type userInfo struct {
//	Id       int64  `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
//	Username string `gorm:"column:username;NOT NULL" json:"username"`
//	Password string `gorm:"column:password;NOT NULL" json:"password"`
//}
//
//var gDb *gorm.DB
//
//func init() {
//	var err error
//	dsn := "root:admin@tcp(localhost:3306)/GoWeb?charset=utf8mb4&parseTime=True&loc=Local"
//	gDb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		log.Fatalln("数据库连接失败：", err.Error())
//	} else {
//		fmt.Println("数据库连接成功")
//	}
//}
//func main() {
//	defer func() {
//		sqlDb, err := gDb.DB()
//		if err != nil {
//			log.Fatalln(err.Error())
//		}
//		err = sqlDb.Close()
//		if err != nil {
//			log.Fatalln("数据库连接断开失败", err.Error())
//		} else {
//			fmt.Println("数据库连接断开成功")
//		}
//	}()
//	r, u := selectOneTest("123", "123")
//	if r.Error != nil {
//		fmt.Println(r.Error)
//	}
//	fmt.Println(u)
//}
//func selectOneTest(users, pwd string) (result *gorm.DB, user userInfo) {
//	user = userInfo{}
//	result = gDb.Where("username = ? AND password = ?", users, pwd).First(&user)
//	return
//}
