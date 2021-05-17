package main

import "fmt"

/**
map 操作
*/

func main() {
	//	判断键值对是否存在
	m1 := map[int]string{1: "foo", 2: "bar"}

	// map通过key获取value，会返回两个值 value type, isPresent boolean
	_, ok1 := m1[1]
	_, ok3 := m1[3]

	// 与 if 搭配使用
	if _, ok2 := m1[2]; ok2 {
		fmt.Println("ok2:", ok2)
	}

	fmt.Println(ok1, ok3)

	//	通过 delete(map map, key string) 来删除key
	m2 := map[string]int{"a": 1}

	delete(m2, "a")

	fmt.Println(m2)
}
