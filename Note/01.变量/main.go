package main

import (
	"errors"
	"fmt"
	"strconv"
)

// 定义全局变量 第一种 (定义全局变量不能用 := )
var q1 = 150

// 定义全局变量 第二种: 类型不一样(自动判断类型)
var (
	q2 = 250
	q3 = "str"
)

/*
	变量的作用域:
	局部变量: 在函数内部声明
	全局变量: 在函数外部声明(首字母大写则在整个程序中都有效)
	如果变量是在一个代码块(for if)中,那么这个变量的作用于在该代码块
	可以理解成在哪个花括号中作用于就在哪
*/

func main() {
	/*
		变量以大写字母开头时,可以被外部包的代码所使用,小写则是不可见的,但是他们在整个包的内部是可见且可用的

		默认值:
		int:    0
		string: ""
		float:  0
		bool: false

		+ 号使用
		左右两边是数值时,则为相加
		左右两边是string时,则为连接
	*/

	// 变量声明 第一种: 指定类型 不赋值则使用默认值
	// var a int = 0
	var a int
	fmt.Println("a =", a)

	// 变量声明 第二种: 根据值自行判断变量类型
	var b = 10.55
	fmt.Println("b =", b)

	// 变量声明 第三种: 省略var(不能是声明过的变量, 自行判断变量类型)
	c := "tom"
	fmt.Println("c =", c)

	// 多变量声明 第一种: 类型一样
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	// 多变量声明 第二种: 类型不一样(自动判断类型)
	var m1, s1, f1 = 100, "dad", 10.5
	fmt.Println(m1, s1, f1)

	// 多变量声明 第三种: 类型不一样(自动判断类型)
	m2, s2, f2 := 105, "mom", "9.9"
	fmt.Println(m2, s2, f2)

	// 多变量声明 补充: := 左边,有至少一个没被声明过
	// err前面声明过了,但是第二行又可以,因为x没有声明过
	err := errors.New("test")
	x, err := strconv.Atoi("testttt")
	fmt.Println(x, err)

	//全局变量输出
	fmt.Println(q1, q2, q3)
}
