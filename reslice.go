package main

import "fmt"

func main() {
	// 切片重组
	ar := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := ar[5:7]

	fmt.Println(len(a)) // 2
	fmt.Println(cap(a)) // 5

	//	a 重新分片
	a = a[0:4]
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
