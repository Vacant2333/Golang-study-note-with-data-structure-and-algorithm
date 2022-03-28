package main

import (
	"fmt"
)

/*
	在Go中,只有数据类型一模一样才能赋值 所以只能通过 &数组名 赋值给指针变量,才代表指针变量指向了这个数组
	和C语言一样,Go中的指针无论是什么类型占用的内存都一样 32位4个字节 64位8个字节
*/

func main() {
	a := 10

	// a的地址: &a
	b := &a
	fmt.Printf("type:%T address:%v \n", b, b)

	// 通过指针b修改a的值
	*b = 15
	fmt.Printf("value:%v address:%v \n", a, b)

	// 通过函数传指针
	changeValue(b)
	fmt.Println(a)

	// 指向数组的指针
	arr := [...]int{10, 20, 30, 40}
	p := &arr
	p[3] = 0
	fmt.Println(arr)
}

func changeValue(a *int) {
	*a = 100
}
