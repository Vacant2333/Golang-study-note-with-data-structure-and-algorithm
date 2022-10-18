package main

import "fmt"

/*
	可以理解成是一个队列 遵循先进先出
	信道是带有类型的管道,可以通过  <- 来发送或接收值
	箭头就是数据流的方向
	ch <- x    (将x发送至信道ch中)
	x := <- ch (从ch接收值并赋予x)
	和interface/slice一样 信道在使用前必须创建(make)
	x := make(chan string)
	默认情况下 发送和接收操作在另一端准备好之前都会被阻塞
*/

// 对一个slice求和
func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// 将结果存入信道中
	ch <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	ch := make(chan int)

	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	// 从管道中取结果
	left, right := <-ch, <-ch

	fmt.Printf("left:%v right:%v \n", left, right)

	// 带缓冲的信道 当为空的时候接收方阻塞,填满后发送方阻塞
	ch2 := make(chan int, 10)
	ch2 <- 10
	ch2 <- 20
	fmt.Println(<-ch2, <-ch2)
	close(ch2)

	// range和close
	// 发送者可通过close关闭一个信道(表示没有需要发送的值了)
	// for i := range ch 会不断从信道接收值,直到被关闭
	ch3 := make(chan int, 10)
	go fibonacci(cap(ch3), ch3)
	for v := range ch3 {
		fmt.Println(v)
	}
}

// 往信道中返回斐波那契数列,n是到多少位
func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		// 把第i位存入信道
		ch <- a
		a, b = b, a+b
	}
	// 关闭信道 这个时候再发送数据会引起panic
	// 检查信道是否关闭: v, ok := <- ch
	// 通常情况无需关闭,只有在必须告诉接收者不再需要接收的时候才有必要关闭,如终止一个range
	close(ch)
}
