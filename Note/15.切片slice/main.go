package main

import "fmt"

/*
	切片是数组的一个引用,因此切片是引用类型 遵守引用传递的机制
	切片的使用和数组类似,遍历/访问/求长度都是一样的
	切片的长度是可以变化的,因此切片是一个可以动态变化的数组

	字符串的底层是[]byte数组,所以字符串也支持切片相关操作
	和数组不同,切片只支持判断是否为nil,不支持== !=判断
	可以通过切片再次生成新的切片,两个切片底层指向统一数组

	对于数组 var a[10]int
	a[0:10] = a[:10] = a[0:] = a[:] 都是等价的

	切片的copy函数在12.内置函数中
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
		切片的拷贝操作,详细在12.内置函数中
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

	// 底层指向同一个数组 改动s1 也会改动原数组ar
	ar := [...]int{10, 20, 30, 40, 50}
	s1 := ar[0:2]
	s2 := ar[0:3]
	s1[1] = 10000
	fmt.Println(ar, s1, s2)

	// 字符串的切片操作
	s := "php is the best language"
	fmt.Println(s[2:10])
	ss := []rune("php是世界上最好的语言")
	fmt.Printf("%c", ss[3:4])

	// 删除元素
	ints := []int{1}
	ints = ints[1:]
	fmt.Println(ints)

	/*
		slice的扩容,长度在[0, 1024]内的话,每次扩容是就容量的两倍,如果旧容量大于等于1024
		最终容量从就容量开始循环增加1/4,直到新容量大于等于申请的容量(貌似不太对,挺复杂的)
	*/
	sss := make([]int, 0)
	old := 0
	for len(sss) < 2048 {
		sss = append(sss, 1)
		if old != cap(sss) {
			fmt.Printf("sss cap[%v]\n", cap(sss))
		}
		old = cap(sss)
	}
}

func test(slice []int) {
	// 注意:函数内可以改切片的元素数据
	// 但是不能改cap和len,必须传指针才可以改
	slice[0] = 10
}
