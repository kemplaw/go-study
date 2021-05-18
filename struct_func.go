package main

import "fmt"

/**
结构体方法
*/

type user struct {
	name string
}

type IntSlice []int

func main() {
	u1 := user{"foo"}
	u1.speak()

	s1 := IntSlice{1, 2, 3}
	fmt.Println(s1.sum())
}

// 结构体函数
// 方法可以通过 (接收者 接收者的类型) 方法名(方法参数) 方法返回值 的形式来创建
// 如果指定了接收者，可以直接通过 接收者.方法名() 来调用
func (u user) speak() {
	fmt.Println(u.name + " speak:")
}

// 非结构体函数
func (i IntSlice) sum() (s int) {
	for _, x := range i {
		s += x
	}

	return
}
