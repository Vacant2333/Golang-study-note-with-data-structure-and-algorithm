package main

import "fmt"

/*
	切片是数组的一个引用,因此切片是引用类型 蹲守引用传递的机制
	切片的使用和数组类似,遍历/访问/求长度都是一样的
	切片的长度是可以变化的,因此切片是一个可以动态变化的数组
*/

func main() {
	var arr = [...]int{1, 2, 3}

	// 声明/定义切片
	var a []int
	var b = []int{1, 2, 3, 4, 5}
	c := make([]int, 2)
	d := arr[1:3]

	// 切片的使用
	fmt.Printf("%T %v %d \n", a, a, cap(a))
	fmt.Printf("%T %v \n", b, b[1:3]) // 输出是2和3 也就是 1 <= 下标 < 3 的元素
	fmt.Println(c, d)

	// 切片的遍历
	// 第一种 常规的for循环遍历
	for a := 0; a < len(b); a++ {
		fmt.Printf("%v |", b[a])
	}
	fmt.Println()
	//第二种 for range
	for i, v := range b {
		fmt.Printf("%v %v | ", i, v)
	}
	fmt.Println()

	// 简写
	// [0:end] => [:end]
	fmt.Println("简写1", b[:2])
	// [start:len(arr)] => [start:]
	fmt.Println("简写2", b[3:])
	// [0:len(arr)] => [:]
	fmt.Println("简写3", b[:])

	/*
		切片的扩容
		append的本质就是对数组的扩容,底层会创建一个新的数组newArr
		然后将原来的元素拷贝到新的数组newArr,然后返回newArr
	*/
	b = append(b, d...)
	b = append(b, -1, -2, -3)
	fmt.Println(b)

	/*
		切片的拷贝操作
		copy(target, src)
	*/
	slice := make([]int, 10)
	copy(slice, d)
	fmt.Println("slice:", slice)

	// 切片是引用传递,具体可见例子
	slice1 := []int{1, 1, 1}
	slice2 := slice1
	test(slice2)
	fmt.Println(slice1, slice2)
}

func test(slice []int) {
	slice[0] = 10
}
