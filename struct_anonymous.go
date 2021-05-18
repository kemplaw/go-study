package main

import "fmt"

/**
结构体的匿名字段 以及 内嵌结构体
*/

type innerS struct {
	in1 int
	in2 int
}

// 如果遇到字段命名冲突的问题，
// 外层的同名字段会覆盖内层的同名字段
// 如果 不同的内层字段结构体 下有 相同的字段名，会编译出错
type outerS struct {
	b      int
	int    "匿名字段int,每种数据类型只能有一个匿名字段"
	innerS "结构体也是一种数据类型，所以也可以作为匿名字段来使用"
}

type A struct {
	ax, ay int "两个字段共享一个类型"
}

func main() {
	o1 := outerS{1, 1, innerS{1, 2}}
	fmt.Println(o1)

	//a := A{1, 2}
}
