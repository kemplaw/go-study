package main

import "fmt"

func main() {
	greeting()
}

func greeting() {
	a1 := 11

	a1Ptr := &a1

	*a1Ptr = 22

	fmt.Println(*a1Ptr)

	fmt.Println("hello func")

	dynamicParams("foo", "bar", "next")

	strList := []string{"foo", "bar"}

	dynamicParams("p", strList...)

	f1()

	deepFunc(10)

	funcParam(20, deepFunc)

	nAdd2 := add2()

	fmt.Println(nAdd2(2))
}

func getNum(input int) (x int) {
	return input
}

// 可变参
func dynamicParams(name string, arg ...string) {
	fmt.Println(name, arg)
}

// defer - 在函数return或者结束之后执行
func f1() {
	fmt.Println("top")
	defer f2()
	fmt.Println("bottom")
}

func f2() {
	fmt.Println("defer print")
}

// 递归
func deepFunc(n int) {
	fmt.Println(n)

	if n > 1 {
		deepFunc(n - 1)
	}
}

// 函数作为参数
func funcParam(y int, f func(int)) {
	f(y)
}

//闭包 - 匿名函数被称为闭包
func add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

// 要注意利用缓存提高计算能力，减少内存占用，加快计算效率
