package main

import (
	"bytes"
	"fmt"
)

// 接口嵌套

type IReadWrite interface {
	Read(b bytes.Buffer) bool

	Write(b bytes.Buffer) bool
}

type ILock interface {
	Lock()
	UnLock()
}

type Files interface {
	IReadWrite
	ILock
	Close()
}

func main() {
	var fileA Files
	fmt.Println(fileA)
}
