package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode    string
	HttpPort   string
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("ini配置文件读取错误:", err.Error())
	}
	LoadServer(file)
	LoadData(file)
}

func LoadServer(file *ini.File) {
	// 获取ini配置文件的分区值，并赋予初始值
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":9090")
}
func LoadData(file *ini.File) {
	// 获取ini配置文件的分区值，并赋予初始值
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("admin")
	DbName = file.Section("database").Key("DbName").MustString("goblog")
}
