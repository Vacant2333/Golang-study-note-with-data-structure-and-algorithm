package main

import "fmt"

// 定义全局变量 第一种
var q1 = 150

// 定义全局变量 第二种: 类型不一样(自动判断类型)
var (
	q2 = 250
	q3 = "str"
)

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

	//多变量声明 第三种: 类型不一样(自动判断类型)
	m2, s2, f2 := 105, "mom", "9.9"
	fmt.Println(m2, s2, f2)

	//全局变量输出
	fmt.Println(q1, q2, q3)
}
