package main

import "fmt"

/*
	Go中的结构体和C中的结构体是一样的,都是用来保存一组 不同类型的数据
	和slice map不同的是,只要定义了结构体变量就可以直接使用

	如果成员中包含slice map,必须make之后才能使用 和变量是一样的
*/

// 定义结构体 方式一: 先定义类型 再定义变量
type node struct {
	data int
	next *node
}

func main() {
	// 相当于 var n = node{}
	var n1 node
	n2 := node{1, nil}
	var n3 = node{10, nil}

	// 访问属性
	n1.data = 10
	fmt.Println(n1, n2, n3)

	// 定义结构体 方式二:定义类型时同时定义变量(匿名结构体)
	var stu = struct {
		name string
		age  int
	}{"na", 20}
	fmt.Println(stu)

	// 复杂结构体 slice map必须make之后使用
	var com = struct {
		m   map[string]int
		sli []int
	}{make(map[string]int, 10), make([]int, 10)}
	com.m["test"] = 10
	fmt.Println(com)

	/*
		结构体之间的比较,同类型的结构体之间,使用指针比较则是比较是否为同一个结构体,不用管是否有不可比较的元素
		如果要比较两个相同结构体之间的内容,必须保证结构体的所有参数都是可比较的类型,比如slice或者map就不行
	*/
	student1 := new(student)
	student2 := new(student)
	// 比较两个student的内容是否相等
	fmt.Println(*student1 == *student2)
	studentInfo1 := new(studentInfo)
	studentInfo2 := new(studentInfo)
	// 因为studentInfo有一个slice结构,所以不能进行比较,只能对他们的指针/地址进行比较
	fmt.Println(studentInfo1 == studentInfo2)
}

type student struct {
	name string
	age  int
}

type studentInfo struct {
	name string
	data []int
}
