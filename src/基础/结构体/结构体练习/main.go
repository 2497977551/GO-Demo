package main

import (
	"fmt"
	"os"
)

// 展示菜单
func openingShow() {
	fmt.Println("-----------------------------欢迎来到学生管理系统------------------------")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员")
	fmt.Println("3.查看所有学员")
	fmt.Println("4.退出")
	fmt.Println("---------------------------------------------------------------------")
}

// 获取用户输入的信息，并传入到student构造函数newStudent中
// 并且返回一个student指针
func getInput() *student {
	var (
		id          int
		name, class string
	)
	fmt.Println("请输入学号：")
	fmt.Scanf("%d\n", &id)
	fmt.Println("请输入姓名：")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入班级:")
	fmt.Scanf("%s\n", &class)
	nst := newStudent(id, name, class)
	return nst
}
func choice(ss *setStudent) {
	var s int
	fmt.Println("请选择：")
	fmt.Scanln(&s)
	if s == 1 {
		fmt.Println("选择的是：添加学员")
		git := getInput()
		ss.addStudent(git)
	} else if s == 2 {
		fmt.Println("选择的是：编辑学员")
		git := getInput()
		ss.UpdStudent(git)
	} else if s == 3 {
		fmt.Println("选择的是：查看所有学员")
		ss.ShowStudent()
	} else if s == 4 {
		fmt.Println("选择的是：退出")
		fmt.Println("再见！")
		os.Exit(0)
	} else {
		fmt.Println("输入错误，系统退出！")
		// os.Exit代表退出程序
		os.Exit(0)
	}

}
func main() {
	s := newSetStudent()
	openingShow()
	for true {
		choice(s)
	}

}
