package main

import "fmt"

//接口类型断言

// 【接口类型】的变量可以包含任何类型的值，需要使用类型断言来检测它的【动态】类型

type IVar interface {
	doSomething()
}

type Person1 struct {
}

func (p Person1) doSomething() {
	fmt.Println("do something")
}

func main() {
	p1 := new(Person1)
	var varI IVar = p1

	// 断言的一般形式， 【接口类型的变量】.(接口类型)
	_, ok := varI.(*Person1)

	fmt.Println(ok)
}
