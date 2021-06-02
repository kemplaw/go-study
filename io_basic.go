package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 读写文件

func main() {
	inputFile, inputError := os.Open("text.txt")
	if inputError != nil {
		fmt.Println("open 错误", inputError)

		return // exit the function on error
	}
	defer inputFile.Close()
	inputReader := bufio.NewReader(inputFile)

	for {
		inputString, readerErr := inputReader.ReadString('\n')

		fmt.Println(inputString)

		// 如果读取内容为空 就跳出循环
		if readerErr == io.EOF {
			return
		}
	}
}
