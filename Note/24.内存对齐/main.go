package main

import (
	"fmt"
	"unsafe"
)

func main() {
	s1 := new(s1)
	s2 := new(s2)
	s3 := new(s3)
	fmt.Println("s1:", unsafe.Sizeof(*s1), "s2:", unsafe.Sizeof(*s2), "s3:", unsafe.Sizeof(*s3))
}

type s1 struct {
	a int8
	c int8
	b int32
}

type s2 struct {
	a int8
	b int32
	c int8
}

type s3 struct {
	a int8
	b int8
	c int8
	d int8
	e int32
}
