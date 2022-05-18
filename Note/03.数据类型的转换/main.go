package main

import (
	"fmt"
	"strconv"
)

/*
	在任何程序中都会存在一些变量有着相同的内部结构,但是却表示完全不同的概念,例如一个
	int类型可以表示索引/时间戳/月份,类型声明语句创建了一个新的类型名称,和现有类型具有
	相同的底层结构,用来分割不同的概念的类型,这样即使他们底层类型相同也是不兼容的!
	type 类型名字 底层类型
	type index  int
*/

func main() {

	/*
		Go在不同类型的变量之间赋值时需要显式转换,不能自动转换
		可以从小范围->大范围,也可以大范围->小范围,但是结果会按溢出处理,不会报错
		基本语法:
			T(v)表达式,将值v转换为类型T
			如 var a float32 = float32(b)
	*/

	// int/float 之间的转换
	var a int64 = 999999
	var b int8
	var c float32 = 10.5
	b = int8(a)
	fmt.Println(b)
	b = int8(c)
	fmt.Println(b)

	// 基本类型转string
	// 方式1 fmt.Sprintf()
	n1 := 10
	n2 := 9.5
	n3 := false
	n4 := 'h'
	s := fmt.Sprintf("%v %v %v %c", n1, n2, n3, n4)
	fmt.Println(s)
	fmt.Println(fmt.Sprintln(n1, n2, n3, n4))

	// 方法2 strconv包(FormatFloat详看文档)
	s2 := strconv.FormatInt(int64(n1), 10) + " " +
		strconv.FormatFloat(n2, 'f', 2, 64) + " " +
		strconv.FormatBool(n3)
	fmt.Println(s2)

	// strconv.Itoa(是strconv.FormatInt(i, 10)的简写)
	// strconv.Atoi(是strconv.ParseInt(s, 10, 0)的简写)

	// string转基本数据类型 仍然使用strconv
	str1 := "100"
	str2 := "12.3450"
	str3 := "false"
	int1, _ := strconv.ParseInt(str1, 10, 32)
	float1, _ := strconv.ParseFloat(str2, 64)
	bool1, _ := strconv.ParseBool(str3)
	fmt.Println(int1, float1, bool1)
}
