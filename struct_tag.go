package main

import (
	"fmt"
	"reflect"
)

/**
携带标签信息的 结构体
*/

type tagType struct {
	f1 bool "这是一个标签"
}

func main() {
	refTag(tagType{true})
}

// 获取tag结构体的类型 此处依赖了反射机制
func refTag(t tagType) {
	ttType := reflect.TypeOf(t.f1)

	fmt.Println(ttType)
}
