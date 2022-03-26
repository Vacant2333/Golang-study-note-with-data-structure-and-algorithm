package main

import "fmt"

func main() {
	// len可用于slice string

	// new 用来分配内存,主要分配值类型,第一个参数是类型,返回值是地址
	a := new(int64)
	*a = 100
	fmt.Printf("new a: %T %v %v \n", a, a, *a)
}
