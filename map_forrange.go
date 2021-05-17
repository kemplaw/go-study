package main

import "fmt"

/**
map for range 用法
*/

func main() {

	m1 := map[string]int{"a": 1, "b": 2}

	// map既不按照key 也不按照 value 的顺序排列
	for key, val := range m1 {
		fmt.Println("key: ", key)
		fmt.Println("val: ", val)
	}

	//	如果只想获取key 可以 for _, val
	// 如果只想获取 val 可以 for key :=range
}
