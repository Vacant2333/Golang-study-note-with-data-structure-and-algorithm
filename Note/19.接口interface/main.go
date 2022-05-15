package main

import "fmt"

/*
	接口和USB插槽很像,它定义某种标准,但是不关心具体实现
	定义接口格式:
	type 接口名称 interface {
		方法声明
	}
	只有实现了接口的所有方法 才算实现了接口 才能用它接受数据
	接口中只能有方法的声明
	接口中也可以包含接口 但是不能包含自己(可以理解成继承)
	空接口可以接受任意类型数据

	接口类型变量可以接受实现了该接口类型的变量
	想要访问变量中的属性,必须将接口类型还原为原始类型
	可以理解成接口只记录类型,要还原之后才可以使用

	空的方法声明(也就是interface{})其实就是泛型
*/

// 定义一个接口
type usb interface {
	up()
}

// 定义两个struct
type typeC struct {
	name   string
	length int
}
type Lighting struct {
	name   string
	length int
}

// 实现接口中的所有方法
func (t typeC) up() {
	fmt.Printf("typec up! name:%v length %v \n", t.name, t.length)
}
func (l Lighting) up() {
	fmt.Printf("linghting up! name:%v length %v \n", l.name, l.length)
}

// 实现一个通过接口定义的公用方法
func work(u usb) {
	u.up()
}

// 接口包含(继承)
type usb3 interface {
	usb
	update()
}

func main() {
	t := typeC{"ttt", 10}
	l := Lighting{"lighting!!!", 20}
	work(t)
	work(l)

	// 定义一个空接口类型变量 可以保存任意类型
	var i interface{}
	i = 123
	fmt.Println(i)
	i = "abc"
	fmt.Println(i)

	i = t
	// 接口类型转换   i.(typeC) 这种格式称之为 类型断言
	// 接口必须转换为原始类型才能访问属性 否则只能调用方法
	if t2, ok := i.(typeC); ok {
		fmt.Printf("t2: %T %v \n", t2, t2)
	}

	var u usb = &l
	// 用type switch转换(不支持fallthrough)
	switch l2 := u.(type) {
	case Lighting:
		fmt.Println(l2)
	}
}
