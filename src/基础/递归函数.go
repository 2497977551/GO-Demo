package main

import "fmt"

func fc(n int) (result int) {
	if n > 0 {
		result = n * fc(n-1)
		return result
	}
	return 1
}

func main() {
	var x int = 2
	fmt.Println(x, fc(x))
}
