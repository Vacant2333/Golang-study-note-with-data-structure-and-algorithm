package main

import (
	"fmt"
	"unsafe"
)

/*
	测试平台:M1 16G
	对齐值取结构体中最长的元素,如在64位下可取的长度是1/2/4/8个字节,
	如果最长的元素是int32(4字节),那么这个结构体就会按4字节对齐.
	尽量对齐,可以节省内存的使用
*/

func main() {
	s1 := new(s1)
	s2 := new(s2)
	s3 := new(s3)
	s4 := new(s4)
	s5 := new(s5)
	fmt.Println("s1:", unsafe.Sizeof(*s1))
	fmt.Println("s2:", unsafe.Sizeof(*s2))
	fmt.Println("s3:", unsafe.Sizeof(*s3))
	fmt.Println("s4:", unsafe.Sizeof(*s4))
	fmt.Println("s5:", unsafe.Sizeof(*s5))
}

// 对齐值:4个字节 长度:8个字节
type s1 struct {
	a int8
	c int8
	// a和c本身长度都是1个字节,但是由于对齐值是4个字节所以他们要占用4个字节
	// 1+1+(浪费2)+4最后加起来就是8个字节
	b int32
}

// 对齐值:4个字节 长度:12个字节
type s2 struct {
	a int8
	b int32
	c int8
}

// 对齐值:4个字节 长度:8个字节
type s3 struct {
	// 不浪费任何空间,4个int8刚好达到对齐值4个字节
	a int8
	b int8
	c int8
	d int8
	e int32
}

// 对齐值:1个字节 长度:2个字节
type s4 struct {
	// 对齐值是1个字节,所以不浪费任何空间
	a int8
	b int8
}

// 对齐值:8个字节 长度:16个字节
type s5 struct {
	a int8
	b int64
}
