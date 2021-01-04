package main

import "fmt"

func main() {
	// 生命一个map集合
	i := make(map[string]string)
	// map通过key-value方式插入键值对
	i["1"] = "Hello,Golang"
	i["2"] = "Hello,Python"
	fmt.Println(i)
	/*使用键输出 */
	for a := range i {
		fmt.Println(i[a])
	}
	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	// 如果key存在ok为true,v为对应的值；不存在ok为false,v为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}
