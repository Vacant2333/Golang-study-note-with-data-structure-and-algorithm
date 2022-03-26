package main

import "fmt"

/*
函数的正常使用:
func 函数名 (形参, a int) (返回值类型1, float32) {
	执行语句...
	return (返回值列表)
}
*/
func test(a int, b int) (int, error) {
	return a + b, nil
}

/*
每一个源文件都可以包含一个init函数,该函数会在main函数执行之前被调用
一般用来做初始化, 类似__construct 折构(构造)函数

如果一个文件同时包含全局变量定义+init+main函数
执行流程: 全局变量定义->init->main
*/

var t int

func init() {
	fmt.Println("init....")
	t = 90
	fmt.Println(t)
}

func main() {
	// 匿名函数
	// 第一种 在定义的时候直接调用 这种方式只能调用一次
	// 求两个数的和
	result := func(n1 int, n2 int) int {
		return n1 + n2
	}(10, 20)
	fmt.Println(result)
	// 第二种 用变量来存放这个函数(a的数据类型是func类型) 这种方式可以反复调用
	a := func(n1 int, n2 int) int {
		return n1 + n2
	}
	fmt.Printf("%T %v\n", a, a(1, 2))

	fmt.Println(fun1(666))
}

// 匿名函数的第三种用法 全局匿名函数
var fun1 = func(n1 int) int {
	return n1 * 10
}
