package main

/*
	连接MySQL
*/
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type test struct {
	Name string
	Sex  string
}

var gOrmDB *gorm.DB

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	var err error
	dsn := "root:admin@tcp(localhost:3306)/GoWeb?charset=utf8mb4&parseTime=True&loc=Local"
	gOrmDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("数据库连接失败：", err.Error())
	} else {
		fmt.Println("连接成功")
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

	//user := test{Name: "Athena", Sex: "女"}
	/*
		相当于insert语句
		INSERT INTO `tests` (`Name`,`Sex`) VALUES ( "Athena", "女")
	*/
	//gOrmDB.Select("ID", "Name", "Sex").Create(&user)

	// 批量插入
	userMax := []test{{"kyo", "男"}, {"JayChou", "男"}, {"JJLin", "男"}}
	// 插入所有字段
	//gOrmDB.Create(&userMax)

	// 批量插入二：指定插入的数量(第二个参数)
	gOrmDB.CreateInBatches(userMax, 3)
}
