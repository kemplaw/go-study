package main

import "fmt"

/**
切片的遍历
*/

func main() {
	var slice1 []int = make([]int, 4)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for idx, val := range slice1 {
		fmt.Println(idx, val)
	}
}
