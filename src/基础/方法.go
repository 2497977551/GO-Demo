package main

import "fmt"

// 方法的实例

// 定义结构体
type person2 struct {
	name string
	age  int
}

// 定义构造函数
func newPerson2(name string, age int) *person2 {
	return &person2{
		name: name,
		age:  age,
	}
}

// 定义方法,该方法属于person2类型.person是接收者
func (p person2) love() {
	fmt.Printf("%s的最爱是编程\n", p.name)
}

// 定义指针接收者，指针是为了数据的一致性，
//当在方法里面修改接收者的值的时候是修改内存地址的值，而不是副本
func (p *person2) setAge(age int) {
	p.age = age
}
func main() {
	//	 调用方法
	p1 := newPerson2("joshua", 20)
	p1.love()

	// 调用方法
	fmt.Println("调用方法前", p1.age)
	p1.setAge(50)
	fmt.Println("调用方法后", p1.age)
}
