package main

/*
	连接MySQL
*/
import (
	"fmt"
	"gorm.io/gorm"
)

type test struct {
	Name string
	Sex  string
}

var gOrmDB *gorm.DB

func init() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	var err error
	err, gOrmDB = IndexInit()
	if err != nil {
		fmt.Println("MySQL数据库连接失败...", err.Error())
	} else {
		fmt.Println("MySQL数据库连接成功")
	}

}
func main() {
	defer IndexDefer(gOrmDB)

	//user := test{Name: "Athena", Sex: "女"}
	/*
		相当于insert语句
		INSERT INTO `tests` (`Name`,`Sex`) VALUES ( "Athena", "女")
	*/
	//gOrmDB.Select("ID", "Name", "Sex").Create(&user)

	// 批量插入
	userMax := []test{{"AMei", "女"}, {"Zed", "男"}, {"Jinx", "女"}}
	// 插入所有字段
	//gOrmDB.Create(&userMax)

	// 批量插入二：指定插入的数量(第二个参数)
	gOrmDB.CreateInBatches(userMax, 3)
}
