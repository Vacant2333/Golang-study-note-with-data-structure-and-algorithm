package main

import "fmt"

/*
	当执行到defer的时候,会将defer后面的语句压入到一个独立的栈中,暂时不执行
	当函数执行完毕(或是发生异常)后,再从栈中取出(先入后出, 也就是从下往上执行)
	也可以理解成是在return之前执行(执行完毕后立马执行return)

	defer推迟调用的函数其参数会立即求值(也就是他的参数不会被后面的语句改变)
*/

func sum(n1 int, n2 int) int {
	defer fmt.Println("n1=", n1) // 3.
	defer fmt.Println("n2=", n2) // 2,
	res := n1 + n2
	// 这一句 不会改变defer的输出 defer会同时把值拷贝存入到栈
	n1 = 0
	fmt.Println("res=", res)
	return res // 1.
}

func main() {
	fmt.Println(sum(1, 5)) // 4.
}
