package main

import "fmt"

/**
map 类型的切片
*/

func main() {
	//	必须使用两次map来初始化一个map类型的切片

	// 通过make 创建切片
	items := make([]map[int]int, 5)

	for i := range items {
		// 通过make 创建map
		// 必须通过 items[i] 设置map的值，如果通过 item ，只是map的值拷贝，并不能真正初始化map的值
		items[i] = make(map[int]int, 1)
		items[i][1] = i + 1
	}

	fmt.Println(items)

}
