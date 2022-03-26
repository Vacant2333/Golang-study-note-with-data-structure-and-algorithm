package main

import (
	"errors"
	"fmt"
)

/*
	在go中处理错误使用defer panic recover
*/

// defer + recover 处理错误
func test1() {
	defer func() {
		// recover()内置函数 可以捕获到异常
		err := recover()
		if err != nil {
			fmt.Println("err=", err)
		}
	}()
	n1 := 10
	n2 := 0
	res := n1 / n2
	fmt.Println("res", res)
}

func main() {
	test1()

	// 自定义错误
	// errors.New(string) 返回一个error类型的值,表示一个错误的说明
	// panic内置函数 接受一个interface{} 输出错误信息并且退出程序(类似php的exit)
	panic(errors.New("test error"))
}
