// 数组与切片
package main

import "fmt"

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
