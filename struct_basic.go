package main

import "fmt"

/**
struct 结构
*/

type IUser struct {
	username string
	age      int
}

// IUserDeep 递归结构体
type IUserDeep struct {
	username string
	age      int
	children *IUserDeep
}

// IStu 结构体转换
type IStu IUser

func main() {
	// 通过new给一个结构体分配内存，返回该结构体在内存中的指针
	user := new(IUser)

	// 惯用的结构体初始化方式
	user1 := &IUser{username: "foo", age: 18}

	// 不提供字段名 初始化时，字段的顺序必须按照结构体的字段顺序来写
	var user2 = IUser{"foo", 18}

	// 结构体可以通过 . 符号来 访问 或者 设置 字段的值
	// 在 Go 中，将 . 符号称为选择器
	user.username = "foo"
	user.age = 10

	fmt.Println(user)
	fmt.Println(user1)
	fmt.Println(user2)

}
