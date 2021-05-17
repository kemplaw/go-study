package main

import "fmt"

/**
将 map 的key 和 value 对调
*/

func main() {
	// 如果map的 value 类型可以作为key，并且是唯一的，就可以与key对调
	// 为了保证不出错，需要仔细检查value的唯一性，并且使用多值map ，例如 map[int][]string

	m1 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("before inverted map: ", m1)

	m2 := make(map[int]string, len(m1))

	for k, v := range m1 {
		m2[v] = k
	}

	fmt.Println("inverted map: ", m2)
}
