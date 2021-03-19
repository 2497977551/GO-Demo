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
	JwtKey     string

	AccessKey  string
	SecretKey  string
	BucKet     string
	QiniuSever string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("ini配置文件读取错误:", err.Error())
	}
	LoadServer(file)
	LoadData(file)
	LoadFile(file)
}

func LoadServer(file *ini.File) {
	// 获取ini配置文件的分区值，并赋予初始值
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":9090")
	JwtKey = file.Section("server").Key("JwtKey").MustString("2021JoshuaGO1017")
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
func LoadFile(file *ini.File) {
	AccessKey = file.Section("qiniu").Key("AccessKey").String()
	SecretKey = file.Section("qiniu").Key("SecretKey").String()
	BucKet = file.Section("qiniu").Key("BucKet").String()
	QiniuSever = file.Section("qiniu").Key("QiniuSever").String()
}
