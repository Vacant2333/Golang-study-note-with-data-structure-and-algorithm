package main

import (
	"fmt"
	"reflect"
)

/*
	类型断言用于检查接口变量的所持有的值是否是期望的数据类型

	动态类型:变量除了有静态类型外(变量声明中的类型),接口变量还有动态类型.
	在程序执行的过程中,其动态类型会随着实现的接口不同,其值也会随之改变.
*/

func main() {
	// 检查是不是int类型 并且转换为int
	var n1 interface{}
	n1 = 50
	// 优雅的处理错误,这里如果ok是false不会panic
	v, ok := n1.(int)
	// 50 true int
	fmt.Println(v, ok, reflect.TypeOf(v))

}
