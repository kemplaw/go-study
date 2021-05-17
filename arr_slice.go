// 数组与切片
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var arr [5]int
	var arr1 = [...]int{1, 2, 3, 4, 5}

	arr[4] = 2

	fmt.Println(arr)

	for i, v := range arr1 {
		fmt.Println(i, v)
	}

	//	数组是值类型，可以用new创建
	//	通过new 创建数组
	var arr3 = new([5]int)

	fmt.Println(arr3)

	arr4 := *arr3

	fmt.Println(arr4)

	var arr5 = [3]int{1, 2, 3}

	f3(arr5)
	f4(&arr5)

	// 数组长度可以被忽略 例如 [5]int [...]string 其中的 5 以及 ... 都可以被忽略
	var arr6 = []int{1, 2, 3}
	fmt.Println(arr6)
	var a7 = []string{3: "foo", 4: "bar"}
	fmt.Println(a7)

	//	多维数组
	var a8 = [][]int{{1, 2}, {1, 2, 34}}

	fmt.Println(a8)

	//	定义切片
	// 切片是引用类型，
	var a9 = [3]int{1, 2, 3} // 普通数组，有长度
	s1 := []int{1, 2, 3}     // 切片，无需长度
	s2 := a9[:2]
	fmt.Println("切片 s1：", cap(s1))
	fmt.Println("切片 s1 len：", len(s1))
	fmt.Println(s2)

	s3 := []int{1, 2, 3, 4, 5}
	s4 := s3[1:3]
	s5 := s4[0:4]

	fmt.Println(s5)

	fmt.Println("通过切片传参实现累加", s6(s3))
	fmt.Println("s4 切片的长度为", cap(s4))

	//	通过 make 创建切片
	var slice1 []int = make([]int, 10)

	fmt.Println(slice1)

	// 关于 len 和 cap 的区别
	// len 是当前切片的长度，是[start, end] start 到 end 的长度
	// cap 是当前切片的最大容量，[start, end] start 到数组结尾的长度
	sliceLenAndCap()

	//	问题1：len 和 cap 分别为
	slice2 := make([]byte, 5)
	fmt.Println("slice2 len", len(slice2))
	fmt.Println("slice2 cap", cap(slice2))

	//	问题2
	slice2 = slice2[2:4]
	fmt.Println("slice2 len", len(slice2))
	fmt.Println("slice2 cap", cap(slice2))

	// 问题3 引用的例子
	slice3 := []byte{'p', 'o', 'e', 'm'}
	slice4 := slice3[2:]

	fmt.Println(slice4)
	fmt.Println("slice3", slice3)
	// 由于引用关系，所以改变了 slice4 其实也改变了 slice3
	slice4[1] = 't'
	fmt.Println(slice3)
	fmt.Println(slice4)

	//	Buffer
	//	通过 buffer 串联字符串
	buffer := bytes.Buffer{}

	for {
		if s, ok := getNextString(); ok {
			buffer.WriteString(s)
			fmt.Println(buffer)
		} else {
			break
		}
	}

}

func getNextString() (s string, ok bool) {
	return "foo", true
}

func f3(a [3]int) {
	fmt.Println(a)
}

func f4(a *[3]int) {
	fmt.Println(a)
}

// 数组作为函数参数
// 如果直接传递数组，会消耗很多内存资源
// 1. 通过指针传递 - 不常用
func f5(ls *[]int) {
	fmt.Println(ls)
}

// 2. 通过数组切片传递
func s6(ls []int) (sum int) {
	for _, v := range ls {
		sum += v
	}

	return sum
}

// 切片 cap 与 len 的比较
func sliceLenAndCap() {
	s1 := [4]int{1, 2, 3, 4}
	s2 := s1[1:3]

	fmt.Println("s2 的 cap 值为：", cap(s2)) // 3
	fmt.Println("s2 的 len 值为：", len(s2)) // 2
}
