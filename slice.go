package main

import (
	"fmt"
	"unicode/utf8"
)

/**
切片
*/

func main() {

	//	切片的复制与追加
	// 如果想增加切片的【容量】，需要创建一个新的 更大容量的切片，然后将原切片的内容都拷贝过来
	// 如何实现：
	// 1. 可以通过 copy 函数，拷贝切片实现这一效果
	// 2. 可以通过 append 函数，向切片追加新元素
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	// 通过copy函数增加切片容量
	n := copy(slTo, slFrom)

	fmt.Println("通过 copy 函数新增切片的结果，sl_to: ", slTo)
	fmt.Println("切片长度 n:", n)

	//	通过 append 函数追加切片
	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println("通过 append 函数新增切片的结果：sl3:", sl3)

	//	从字符串生成字节切片
	// 字符串实质上是一个字节的切片
	s := "foo"
	c := []byte(s)
	fmt.Println("s字符串的字节切片结果:", c)

	//	计算字符串中字符的数量
	//	1. 通过 len([]int32(s)) 来计算，但是效率不如 utf8.RuneCountInString(s) 高
	fmt.Println("字符串s的字符数量为：", utf8.RuneCountInString(s))

	//	获取字符串的一部分 str[start:end]
	//  str[start:] 表示从start到len(str) - 1 位置的字符串
	// str[:end] 表示从 0 开始到 end-1 处的字符串
	substr := s[0:1]
	fmt.Println("substr: ", substr)

	//	字符串与切片的内存结构
	//	字符串实际是一个双字结构，一个指向实际数据的指针 和 一个记录字符串长度的整数

	//	如何修改字符串中的字符
	// 字符串无法被直接修改，如 str[1] = 'd' 这样的操作是错误的
	//	可以通过 字符串 -> 字节数组 -> 操作字节数组 -> 转为字符串 这样的流程来达到修改字符串的目的
	// 此处可以封装一个 strReplace(index int, val byte) string 函数，来达到修改字符串的目的
	c[0] = 'b'
	s2 := string(c)
	fmt.Println("修改之后的字符串s为：", s2)

	//	搜索及排序切片和数组
	//	只有排序之后的切片才能被搜索，因为标准库的搜索算法使用的是二分法
	//	int类型的切片可以使用 sort.Ints()，其他类型对应其他类型的sort方法
	//	更多参考官方文档 http://golang.org/pkg/sort/

	//	append 操作切片的常用操作

	//	1. 切片b追加到切片a之后
	var slB = []int{1, 2, 3}
	var slA = []int{0, 1, 2}
	slC := append(slA, slB...)
	fmt.Println(slC)

	//	其他操作详见 https://www.bookstack.cn/read/the-way-to-go_ZH_CN/eBook-07.6.md

	//	切片和垃圾回收
	//	切片底层其实是一个【数组】，【数组实际容量】可能要大于【切片定义的容量】，只有【不存在任何切片指向】时，数组内存才会被释放
	//	这个特性可能会导致程序占用多余内存，例如

	//	1. 读取某个文件，尝试获取其中的一部分内容，并返回这部分内容的切片。
	//	这容易导致，这部分内容的切片内存如果不释放，可能这个文件的内存也不会被回收，如何改进呢？

	//	改进方案：可以通过拷贝需要的部分到一个新的切片中，避免这样的问题
	// c := make(byte[], len(file))
	// copy(c, file)
}
