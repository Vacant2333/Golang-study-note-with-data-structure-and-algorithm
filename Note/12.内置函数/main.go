package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
		len 可用于 slice array string
		cap 可用于 slice array
		cap 是获得容量,len是获得长度,slice的容量一般为长度*2

		make(type, len, cap) 用于slice/map/chan

		append在 15.切片slice 有解释
	*/

	/*
			new 用来分配内存,主要分配值类型,第一个参数是类型,返回值是地址
			注意: 如果类型的大小是0,例如struct{} [0]int,有可能返回相同的地址(依赖具体的实现)
		          谨慎使用大小为0的类型
			new函数使用相对比较少,一般用字面量语法创建会更灵活  student{"xx"}
	*/
	a := new(int64)
	*a = 100
	fmt.Printf("new a: %T %v %v \n", a, a, *a)

	// 大小为0的类型
	p1 := new([0]int)
	p2 := &struct{}{}
	fmt.Println("[0]int sizeof: ", unsafe.Sizeof(*p1), " p1: ", p1)
	fmt.Println("struct {}{} sizeof: ", unsafe.Sizeof(*p2), "p2: ", p2)

	// new只是一个预定义的函数,并不是一个关键字,因此我们可以将new定义为别的类型
	new := 10
	fmt.Println(new)
}
