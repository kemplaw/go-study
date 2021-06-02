package main

import "log"

// 一个类型 内嵌 另一个的指针时，就可以使用另一个类型所有的接口方法

type Task struct {
	Command string
	*log.Logger
}

func main() {
}
