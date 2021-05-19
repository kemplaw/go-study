package main

import (
	"./person"
	"fmt"
)

/**
结构体方法
*/

type user struct {
	name string
}

type IntSlice []int

//  内嵌类型的方法与继承
//  如果在A类型中包含B类型，则A类型继承了B类型的方法和字段

// 继承包括 匿名继承 以及 具名继承
// 匿名继承，可以直接使用继承到的字段 如 car.Start()
// 具名继承，需要通过字段名来使用继承到的字段，如 car.Engine.Start()

// 多重继承：可以在结构中内嵌多个其他的结构类型达到多重继承的目的
/**
type A struct {
	B
	C
}
*/

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
}

func main() {
	u1 := user{"foo"}
	u1.speak()

	s1 := IntSlice{1, 2, 3}
	fmt.Println(s1.sum())

	u2 := new(user)
	u2.name = "bar"
	u2.changeName("abc")

	// 由于 person包没有导出 name 以及 age，
	// 直接设置会出现异常 Cannot assign a value to the unexported field 'xxx'
	//p1 := person.Person{"foo", 11}

	// 可以通过 oop 中的 getter 和 setter 概念来操作
	p1 := new(person.Person)

	p1.SetName("foo")
	fmt.Println(p1.Name())

	// 实现继承
	car := new(Car)
	car.Start()
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

// 函数与方法之间的【区别】
// 函数 将变量作为参数
// 函数的参数是指针时，也可以改变参数的值或者状态

// 方法 是在变量基础上进行调用 如 val.func()
// 方法在接收者是指针时，可以改变接收者的值或者状态
// 接收者相关联的方法不写在类型结构里面，耦合更宽松，数据与行为（方法）是独立的

// 指针作为接收者, 而非值作为接收者，这样可以减少内存占用
func (u *user) changeName(name string) {
	u.name = name

	fmt.Println(u.name)
}
