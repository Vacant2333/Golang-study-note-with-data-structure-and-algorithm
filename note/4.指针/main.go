package main

import (
	"fmt"
)

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
}

func changeValue(a *int) {
	*a = 100
}
