package main

import "fmt"

/*
	空接口
	空接口就是没有定义任何方法的接口，任何类型都可以实现空接口
	空接口可以接收任意值
*/
func main() {
	/*
		定义一个空接口类型的变量
		可以看出空接口可以接收任意类型的值
	*/
	var x interface{}
	x = 18
	fmt.Println(x)
	x = "joshua"
	fmt.Println(x)
	x = true
	fmt.Println(x)

	//	空接口也可以运用到map上
	var i = make(map[string]interface{}, 20)
	i["name"] = "Athena"
	i["age"] = 16
	i["hobby"] = []string{"code", "sleep", "LOL"} // string类型的切片
	fmt.Println(i)

	/*
		类型断言： x.(T)
		x：表示类型为interface{}的变量
		T：表示断言x可能是的类型。
		该语法返回两个参数，第一个参数是x转化为T类型后的变量，
		第二个值是一个布尔值，若为true则表示断言成功，为false则表示断言失败
	*/
	// if方法断言
	res, ok := x.(string)
	if !ok {
		fmt.Println("类型不是string", ok)
	} else {
		fmt.Println(res)
	}

	//	switch方法断言
	switch v := x.(type) {
	case string:
		fmt.Printf("x是string类型，结果是：%v\n", v)
	case bool:
		fmt.Printf("x是bool类型，结果是：%v\n", v)
	case int:
		fmt.Printf("x是int类型，结果是：%v\n", v)
	case float64:
		fmt.Printf("x是float64类型，结果是：%v\n", v)
	}
}
