package main

import "fmt"

/*
指针
*/

func main() {
	i1 := 5

	fmt.Println("i1的指针为（可能每次运行都不一样）：", &i1)

	// 一个指针变量可以指向任何一个值的内存地址，32位机器占用 4字节 64位8字节
	// 定义指针之后，没有分配至任何变量时，默认值为nil
	//	指针地址可以保存在【指针】这个特殊类型中，例如
	var intP *int = &i1

	fmt.Println(intP)

	//	一个指针变量一般缩写为ptr
	name := "foo"
	var namePtr *string = &name

	fmt.Println(namePtr)

	//	*号放置在指针前，可以得到该指针对应的值（反引用或者间接引用），如下：
	var namePtrVal string = *namePtr
	fmt.Println(namePtrVal)

	// 通过对 间接引用 的值赋值，可以改变该引用值的值
	namePtrVal = "bar"

	fmt.Println(namePtrVal)

	//	无法获取【字面量】或者【常量】的地址
	//	plt := &"foo" // error

	//  const name = "bar"
	//	plt := &name // error

	//  在 Go 中，对指针进行加减操作是不合法的

	//	比较优雅的用法
	//	1. 在函数中，通过指针来传递参数，达到优化性能的目的（这是因为，函数的参数一般使用按值传递，会拷贝一份参数的值，通过指针可以减少这样的内存资源）
	//	2. 不建议指针指向另一个指针，会造成逻辑的嵌套，不易于维护
}
