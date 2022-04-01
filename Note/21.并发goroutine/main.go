package main

import (
	"fmt"
	"time"
)

/*
	goroutine是由Go运行时管理的轻量级线程
	go f(a, b, c)
	a,b,c的求值在当前进程中,而执行发生在新的进程中
*/

func say(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}

func main() {
	for x := 0; x < 10; x++ {
		go say(fmt.Sprintf("我是: %v", x))
	}
	// 挂起主进程 否则主进程结束后其他进程会结束
	for {

	}
}
