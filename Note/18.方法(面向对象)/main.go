package main

import "fmt"

/*
	Go中的方法其实就是一个特殊的函数,只是这个函数是和某种属性类型绑定在一起的而已
	方法一般用于将函数和结构体绑定在一起,让结构体除了能够保存数据外还能具备某些行为
	方法的数据类型也是函数类型,所以也可以定义变量保存(作为返回值等)

	将函数与数据类型绑定的格式:
	func (接受者 数据类型) 方法名称 (形参列表) (返回值列表) {
		方法体...
	}
*/

type stu struct {
	name string
	age  int
}

// 定义一个和stu绑定的方法
func (s stu) say() {
	fmt.Printf("say name:%v age:%d \n", s.name, s.age)
}

// 接受者是一个指针(这样就是引用传递,内部修改会影响方法的外部)
func (s *stu) say2() {
	fmt.Printf("say2 name:%v age:%d \n", s.name, s.age)
}

func main() {
	s := stu{"小明", 95}
	s.say()
	// 也可以直接 s.say2() 底层会自动获取地址传递给接受者
	(&s).say2()
}
