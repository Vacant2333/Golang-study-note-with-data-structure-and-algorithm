package main

import "fmt"

/*
	数组一旦声明了长度就是固定的 不能动态变化
	数组中的元素可以是任何数据类型 包括值类型和引用类型
	创建后如果没有赋值是有默认值的 0/""/false
*/

func main() {
	// 定义数组(不加...是切片)
	var a [6]int
	var b = [3]int{1, 2, 3}
	var c = [...]int{4, 5, 6}
	var d = [...]string{1: "b", 0: "a", 2: "c"}

	// 赋值
	a[0] = 10

	// 数组的使用
	fmt.Println(a[0], a[1])
	fmt.Printf("%T %v \n", b, b)
	fmt.Printf("%T %v \n", c, c)
	fmt.Printf("%T %v \n", d, d)

	// 数组的遍历
	for index, value := range d {
		fmt.Printf("d[%d] : %T %v \n", index, value, value)
	}

	// 通过函数修改数组的值
	test(&a)
	fmt.Println(a)
}

func test(p *[6]int) {
	(*p)[0] = 12345
}
