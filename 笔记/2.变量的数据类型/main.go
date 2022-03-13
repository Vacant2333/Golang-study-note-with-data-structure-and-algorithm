package main

import (
	"fmt"
	"unsafe"
)

func main() {
	/*
		Go在整形变量使用中, 遵守保小不保大的原则
		即: 在保证程序正确运行的情况下,尽量使用占用空间小的数据类型

		int/float后面的数字是多少个bit,8bit = 1byte,uint是非负数类型
		byte: uint8的别名
		rune: int32的别名(表示一个Unicode码点,可以表示中文)

		通常使用float64,比float32更精确,声明时默认float64
		float中 .512 = 0.512

		反引号用来创建原生的字符串字面量,不支持任何转义,可以由多行组成,一般用于多行消息,HTML,正则表达式
		在Go中字符串是不可变的,一旦赋值就不能修改

		科学计数法:
		5.12e2 = 5.12 * 10^2
		5.12e-2 = 5.12 / 10^-2

		基本数据类型:
			数值型:
				整数类型: byte, int, int(8/16/32/64), uint, uint(8/16/32/64)
				浮点类型: float(32/64)
			字符型:
				没有专门的字符型, 一般用byte保存单个字母/字符,不能存汉字(go中编码用utf-8,汉字占2个字节)
			布尔型:
				bool (false/true)
			字符串:
				string
			复数:
				complex32, complex64
		派生/复杂数据类型:
			指针(Pointer)
			数组
			结构体(struct)
			管道(chan)
			函数(也是一种数据类型)
			切片(slice)
			接口(interface)
			map
	*/

	// byte, rune 的使用
	var a byte
	var b rune
	a, b = 'a', '是'
	fmt.Println(a, b)
	a = 90
	fmt.Printf("%c %c\n", a, b)

	// float 的使用
	var c float32
	c = 1.23456789
	fmt.Println(c)

	// int 存字符
	var d int = '中'
	fmt.Printf("%c %d\n", d, d)

	// 查看变量的数据类型和占用的字节数
	var demo uint64
	fmt.Printf("数据类型: %T 占用的字节数: %d", demo, unsafe.Sizeof(demo))
}
