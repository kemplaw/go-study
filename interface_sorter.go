package main

import "fmt"

// 使用 Sorter 接口排序

type Sorter interface {
	len() int
	less(i, j int) bool
	swap(i, j int)
}

// IntArray 声明数组
type IntArray []int

func (arr IntArray) len() int {
	return len(arr)
}

func (arr IntArray) less(i, j int) bool {
	return arr[i] < arr[j]
}

// 交换数组中值的位置
func (arr IntArray) swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	// 转型
	a := IntArray(data)

	Sort(a)

	fmt.Println(a)
}

func Sort(data Sorter) {
	for pass := 1; pass < data.len(); pass++ {
		for i := 0; i < data.len()-pass; i++ {
			if data.less(i+1, i) {
				data.swap(i, i+1)
			}
		}
	}
}
