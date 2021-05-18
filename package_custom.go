package main

import (
	"./pack1"
	"fmt"
)

/**
自定义包
*/

func main() {
	fmt.Println(pack1.ReturnStr())

	//	如果通过 import . from 'path/to' 可以直接使用该包中的项目
	//	如果通过 import _ from 'path/to' 只导入副作用，只会执行包中的 init 函数，并初始化包中的变量

	//	如果需要使用【外部包】，需要使用 go install 包的地址，如 go install codesite.ext/author/goExample/goex
	//	或者一次性使用的方式 import goex "codesite.ext/author/goExample/goex"
}
