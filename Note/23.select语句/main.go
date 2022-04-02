package main

import (
	"fmt"
	"time"
)

/*
	select是Go中的一个控制结构,类似switch,用于处理异步IO操作
	如果有多个case可以运行,那么会随机选出一个运行
	如果没有可运行的case,但是有default,那么就会执行default
	如果没有可运行的case,且没有default,那么就会阻塞,直到某个case可以运行
	一般情况下使用select不用写default
	一般情况会用for嵌套,否则select只执行一次
	应用场景: 多路监听/超时处理
*/

func main() {
	var ch = make(chan int)
	var exitCh = make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			// 向管道中添加数据
			ch <- i
			time.Sleep(time.Second)
		}
		exitCh <- true
	}()

	// 读取数据
	flag := true
	for flag {
		fmt.Println("开始读取管道:")
		select {
		case data := <-ch:
			fmt.Println("读到了:", data)
		case <-exitCh:
			fmt.Println("exitCh读到了,退出")
			flag = false
		}
	}

	// 超时处理
	// time.After 会在指定时间后向对应信道传入一个time类型数据
	select {
	case <-ch:
		fmt.Println("不会到这里,ch没数据了")
	case t := <-time.After(2 * time.Second):
		fmt.Println("超时了 t:", t)
	}
}
