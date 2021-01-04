package main

import "fmt"

func main() {
	index := [...]int{3, 4, 6, 1, 7}
	num := len(index)
	for i := 0; i < num; i++ {
		for j := i; j < num; j++ {
			if index[i] < index[j] {
				temp := index[i]
				index[i] = index[j]
				index[j] = temp

			}

		}

	}
	fmt.Println(index)
}
