package main

import "fmt"

func main() {
	i1 := 50
	if i1 < 90 {
		fmt.Println(i1)
	}

	i2 := 10.5
	if i2-10 > 0 {
		fmt.Println("flag1")
	} else {
		fmt.Println("flag2")
	}

	i3 := 10
	switch i3 {
	case 1:
		fmt.Println("one expression")
	case 2, 3, 4, 5: // 多个表达式
		fmt.Println("multiple expression")
	case 10:
		fmt.Println("this one")
	default: //默认
		fmt.Println("default")
	}

	var i4 byte = 'a'
	// switch中只要是个值就行
	switch test(i4) + 1 {
	case 'a':
		fmt.Println("a")
	case 'b':
		fmt.Println("b")
	case 'c':
		fmt.Println("c")
		//switch穿透,会继续执行下一个case 如果匹配到这个case
		fallthrough
	case byte(i3):
		fmt.Println("i3")
	}

	// 不写表达式的switch 类似if else来使用
	switch {
	case i1 == 50:
		fmt.Println("i1")
	case i1 < 0:
		fmt.Println("no")
	case i1 == int(i2):
		fmt.Println("no")
	}

	// switch内声明/定义一个变量 分号结束
	switch grade := 90; {
	case grade > 80:
		fmt.Println("grade")
	}

	// switch判断类型 .(type) 必须用于switch 注意:不能用fallthrough
	var x interface{}
	var y float64 = 3.5
	x = y
	switch x.(type) {
	case float64:
		fmt.Println("float64")
	case int:
		fmt.Println("int")
	case nil:
		fmt.Println("type:%T", x)
	}

	// for 循环
	for z := 1; z <= 10; z++ {
		fmt.Println("z", z)
	}

	// for 只写条件
	z2 := 5
	for z2 < 10 {
		z2++
		fmt.Println("z2", z2)
	}

	// for 不写条件 无限循环
	for {
		fmt.Println("break")
		break
	}

	// for range
	// 不能用传统方式遍历有中文的string,但是可以转成切片来遍历
	s1 := "abcde"
	for index, val := range s1 {
		fmt.Println(index, val)
	}

	// break 通过标签终止一层语句块 默认退出最近的循环
label1:
	for {
		for {
			break label1
		}
	}

	// continue 通过标签跳出一次循环 默认跳出最近的循环
label2:
	for iii := 0; iii < 4; iii++ {
		for iii < 4 {
			// 在iii为2的时候直接continue了大循环 跳到label2的位置继续执行
			if iii == 2 {
				continue label2
			}
			fmt.Println("continue iii=", iii)
			break
		}
	}
}

func test(a byte) byte {
	return a + 1
}
