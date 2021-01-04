package main

import "fmt"

/*
	Go语言的结构体没有构造函数，我们可以自己实现。
	例如，下方的代码就实现了一个person的构造函数。
	因为struct是值类型，如果结构体比较复杂的话，
	值拷贝性能开销会比较大，所以该构造函数返回的是结构体指针类型。
*/
type person struct {
	name, sex string
	age       int
}

func newPerson(name, sex string, age int) *person {
	return &person{
		name: name,
		sex:  sex,
		age:  age,
	}
}
func main() {
	p := newPerson("joshua", "男", 20)
	fmt.Println(p.name)
}
