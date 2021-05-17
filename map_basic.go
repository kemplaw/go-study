package main

import "fmt"

/**
map 基础
*/

func main() {
	// map 是引用类型
	// map的容量是可变的
	//	声明形式：var m1 map[keyType]valueType
	var m1 map[string]int
	// 可以使用make创建map
	m2 := make(map[string]int)

	var m3 map[string]int

	// 赋值方式
	m1 = map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m1)

	//	通过访问key的方式赋值
	m2["foo"] = 1
	m2["bar"] = 2

	fmt.Println(m2)

	// 由于map为引用类型，m3与m1指向同一个引用地址，所以修改了m3，m1也同步修改了。
	m3 = m1
	m3["kemp"] = 1

	fmt.Println(m3, m1)

	//	map 中的 value可以为任意值
	m4 := map[string]func() int{
		"a": func() int {
			return 1
		},
		"b": func() int {
			return 2
		},
	}

	fmt.Println(m4["a"]())

	//	用切片作为map的值
	//	一个key对应多个value，此时可以使用切片作为value
	s1 := []int{1, 2, 3}
	m5 := map[int]*[]int{
		1: &s1,
	}

	fmt.Println(m5)
	fmt.Println(m5[1])
}
