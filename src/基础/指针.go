package main

import "fmt"

func main() {
	i := 66
	var r *int
	fmt.Printf("r的类型为%T\n", r)
	r = &i
	*r = 85
	fmt.Println(*r, r)
	fmt.Println(i, &i)
}
