package main

import "fmt"

// 类型判断
// 可以通过 type-switch 的形式来判断接口变量的动态类型

type Circle1 struct {
}

type Circle2 struct {
}

type Shaper1 interface {
	area() int
}

func main() {
	//	先声明接口变量
	var s1 Shaper1

	// 声明结构体
	c1 := new(Circle1)
	c2 := new(Circle2)

	s1 = c1

	switch s1.(type) {
	case *Circle2:
		fmt.Println("circle2")
		s1 = c2
	case *Circle1:
		fmt.Println("circle1")
	}
}

// 实现接口中定义的方法
func (c1 *Circle1) area() int {
	return 5 * 5
}

func (c2 Circle2) area() int {
	return 12 * 12
}
