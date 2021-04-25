package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("test")

	// 使用条件语句来验证函数输出
	if t, ok := mySqrt(5); ok {
		fmt.Println(t)
	} else {
		fmt.Println("wrong param")
	}

	// switch 分支可以支持表达式
	// 分支不需要使用 break 来结束，该分支代码运行结束会自动跳出 switch
	//	如果需要多个分支执行一部分代码，可以使用 fallthrough
	switch n1 := 0; n1 {
	case 0:
		fallthrough
	case 1:
		fmt.Println("true")
	}

	//	迭代，go 中只有 for 循环
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//	使用条件循环
	i2 := 5

	for i2 >= 0 {
		i2 -= 1
		fmt.Println(i2)
	}

	//	for range 循环，可以用来遍历 数组以及 Map，类似于 for of
	str1 := "foo bar"
	for pos, char := range str1 {
		fmt.Println("pos: ", pos)
		fmt.Println("char: ", char)
	}

	for i := 0; i < 5; i++ {
		var v int
		fmt.Println(v)
		v = 5
	}
}

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	}

	return math.Sqrt(f), true
}
