package main

import "fmt"

// 方法集与接口

type List []int

func (l List) len() int {
	return len(l)
}

func (l *List) Append(val int) {
	// 通过 *指针 的形式来给指针指向的值 赋值
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

// CountInto
// start, end int 表示 start int, end int 的简写
func CountInto(appender Appender, start, end int) {
	for i := start; i <= end; i++ {
		appender.Append(i)
	}
}

type Lener interface {
	len() int
}

func LongEnough(lener Lener) bool {
	return lener.len()*10 > 42
}

func main() {
	//	声明一个 list 类型的变量
	var ls List

	//CountInto(ls, 1, 10) // 此处报错的原因是因为，ls 为list类型，并非 *list 类型，Appender 的 append 方法是定义在指针类型上的

	if LongEnough(ls) {
		fmt.Println("ls is long enough")
	}

	lsPtr := new(List)

	CountInto(lsPtr, 1, 10)

	fmt.Println(lsPtr)
}
