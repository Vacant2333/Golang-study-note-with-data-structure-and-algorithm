package main

import (
	"fmt"
	"reflect"
)

/*
	类型断言用于检查接口变量的所持有的值是否是期望的数据类型

	语法: 左边必须是接口类型 type可以是类型标识符,也可以是类型字面量
	PrimaryExpression.(Type)

	动态类型:变量除了有静态类型外(变量声明中的类型),接口变量还有动态类型.
	在程序执行的过程中,其动态类型会随着实现的接口不同,其值也会随之改变.
*/

type structA struct {
	name string
}

type structB struct {
	name string
}

func (sb structB) run() {
	fmt.Println("sb.run~")
}

func main() {
	// 检查是不是int类型 并且转换为int
	var n1 interface{}
	n1 = 50
	// 优雅地处理错误,这里如果ok是false不会panic
	// 当断言不成立时，第一个值将会作为测试类型的零值
	v, ok := n1.(int)
	// 50 true int
	fmt.Println(v, ok, reflect.TypeOf(v))

	// 动态类型
	var n2 interface{}
	// n2的动态类型是structA
	n2 = structA{"sa"}
	// n2的动态类型是structB
	n2 = structB{"sb"}
	fmt.Println(n2, reflect.TypeOf(n2))

	// type是类型字面量
	var n3 interface{}
	n3 = structB{"sb"}
	fmt.Println(n3.(interface {
		run()
	}))
	n3.(interface {
		run()
	}).run()
}
