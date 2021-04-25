package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IType = int

// 批量定义类型
type (
	ITypeA int
	ITypeB string
	ITypeC float32
)

// 常量 - 赋值时，必须为编译时就能确定的值
const constValue = 1

// 声明变量 - 简单类型可省略
// 变量可以并行赋值
// _ 符号可以用于抛弃值，是一个只写变量，无法获取值
var v = 1

// 批量声明变量
var (
	v1 = 1
	v2 = 2
)

// float 类型尽可能使用 float64

func main() {
	var a IType = 5
	var aCoverToInt int = int(a)
	// 简化声明变量 - 只能在函数体内使用这种声明方式
	v3 := 3

	fmt.Println("hello world")
	fmt.Println("custom type value: ", a)
	fmt.Println("covered type value: ", aCoverToInt)
	fmt.Println(v3)

	var str string = "foo bar"

	fmt.Println("str 字符串是否以 f 为开头：", strings.HasPrefix(str, "f"))
	fmt.Println("str 字符串是否包含 foo：", strings.Contains(str, "foo"))
	fmt.Println("str 字符串中 foo 的index：", strings.Index(str, "foo"))
	fmt.Println("替换 str 字符串中的foo 为 hello：", strings.Replace(str, "foo", "hello", -1))
	fmt.Println("字符串 str 中 o 的出现次数为：", strings.Count(str, "o"))
	fmt.Println("字符串 str 变为大写：", strings.ToUpper(str))
	fmt.Println("str 按空格分隔", strings.Split(str, " ")) // 如果只是空格，可以用 strings.Fields(s) 来简单分割
	fmt.Println("str 按逗号拼接：", strings.Join(strings.Split(str, " "), ","))

	// 字符串转为 int
	var orig string = "666"
	var an, _ = strconv.Atoi(orig)

	fmt.Println("字符串 转为 int 之后的值为：", an)

	// 指针
	// 需要注意 无法获取字面量或常量的地址
	// 可以通过传递指针来 降低程序的内存占用，以及提高效率
	var i1 int = 5
	const c1 int = 5
	var intP *int = &i1
	//var sP *string = &"foo bar" // 错误的使用方法
	//var i2P *int = &c1 // 错误

	fmt.Println("i1 的指针为：", intP)

}
