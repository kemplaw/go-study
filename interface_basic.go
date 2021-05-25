package main

import "fmt"

/**
接口基础
*/

// 接口定义了一组方法（没有具体实现），来说明对象的一组行为
// 接口里不能包含变量

// 多类型可以实现同一个接口
// 类型会隐式的实现接口，不需要显示指定接口

// Shaper 只有一个方法的接口，一般以 er 结尾
type Shaper interface {
	Area() float32
}

// IDoSomething 当er 不合适时，使用 able 或者 I开头的命名方式
type IDoSomething interface {
	doGet()
	doPost()
}

type Square struct {
	side float32
}

func (s *Square) Area() float32 {
	return s.side * s.side
}

func main() {
	// 创建结构体的实例
	sq1 := new(Square)
	sq1.side = 5.25

	// 声明一个接口类型的变量
	var areaInfo Shaper
	// 由于 sq1 实现了 shaper 的 area 方法，所以可以视为 sq1 实现（隐式）了接口 Shaper
	// 所以 areaInfo 可以赋值为 sq1
	// sq1 没有实现 shaper 接口，会抛出 Square not implement interface 的异常
	areaInfo = sq1

	fmt.Println(areaInfo.Area())
}
