package main

import (
	"fmt"
	"sort"
)

/**
map 的排序
*/

func main() {
	//	map 默认是无序的，无论按照key还是value都不会排序
	//	如果想要排序，需要将key或者value拷贝到一个切片，通过sort进行排序

	var m1 map[int]int = map[int]int{0: 1, 1: 1, 2: 2}

	//	将key拷贝到切片
	keys := make([]int, len(m1))

	i := 0

	for key, _ := range m1 {
		keys[i] = key
		i++
	}

	sort.Ints(keys)
	fmt.Println("keys: ", keys)

	for _, val := range keys {
		fmt.Println("key: ", val, keys[val])
	}

	//	但是如果你想要一个排序的列表你最好使用结构体切片，这样会更有效：
	type name struct {
		key   string
		value int
	}
}
