package main

import "fmt"

/**
结构体工厂
*/

type File struct {
	name string
	path string
}

func main() {
	file := NewFile("demo.txt")

	fmt.Println(*file)
}

// NewFile 结构体工厂的命名 一般以 new 或者 New
// 简易工厂函数的实现
func NewFile(name string) *File {
	return &File{name, "path/to/path"}
}
