package main

import "fmt"

func main() {
	/*
		len 可用于 slice array string
		cap 可用于 slice array
		cap 是获得容量,len是获得长度,slice的容量一般为长度*2

		make(type, len, cap) 用于slice/map/chan

		append在 15.切片slice 有解释
	*/

	// new 用来分配内存,主要分配值类型,第一个参数是类型,返回值是地址
	a := new(int64)
	*a = 100
	fmt.Printf("new a: %T %v %v \n", a, a, *a)
}
