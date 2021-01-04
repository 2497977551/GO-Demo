package main

import (
	"encoding/json"
	"fmt"
)

/*
	字段可见性即首字母大写
	如果字段没有对外暴露，是无法转换为json的，
	因为json不属于当前main包，而是go内置的方法

	tag标签：
	1.在字段后面添加反引号并且是键值对格式
	2.值必须用双引号包裹
	3.多个tag可以使用空格隔开
	下方使用的tag表示当使用json方法序列化时可以将他原本的key
	改为tag内定义的键值对中的value作为key
	可以运行查看结果
*/

// 定义学生结构体
type student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// student构造方法
func newStudent(id int, name string) *student {
	return &student{
		ID:   id,
		Name: name,
	}
}

// 定义班级结构体
type class struct {
	Title    string    `json:"title"`
	Students []student `json:"student_list"`
}

func main() {
	c := class{
		Title: "终极一班",
		// 初始化切片
		Students: make([]student, 0, 10),
	}
	// 循环添加学生name以及学生id
	for i := 0; i < 10; i++ {
		// 给构造方法循环添加学生name以及学生id
		sdt := newStudent(i, fmt.Sprintf("学生%02d", i))
		// 将构造方法添加到students切片中
		c.Students = append(c.Students, *sdt)
	}
	fmt.Printf("%T", c)
	fmt.Println(c)

	//	json序列化，转为json格式
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json 转换失败:", err)
	}
	// 输出json 必须要用格式化输出
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)

	//	json 反序列化，将json字符串转为go语言识别的类型
	var c1 class // 定义json反序列化接收的返回值

	jsonStr := `{"Title":"终极一班","Students":[{"ID":0,"Name":"joshua"},{"ID":1,"Name":"kyo"}]}`
	// json.Unmarshal第一个参数为json但必须转为byte切片，
	//第二个参数必须是指针地址不然传不进去，因为是变量是值拷贝
	err1 := json.Unmarshal([]byte(jsonStr), &c1)
	if err1 != nil {
		fmt.Println("json 反序列化失败：", err1)
	}
	fmt.Printf("%v", c1)
}
