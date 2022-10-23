package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.Cond{L: &sync.Mutex{}}

	go func() {
		for i := 0; i < 2; i++ {
			cond.L.Lock()
			// wait时会block住,并且必须要Lock
			// 可以用循环套住wait,如果条件不成立则继续wait
			cond.Wait()
			cond.L.Unlock()
			fmt.Println(i)
		}

	}()
	// 唤醒在wait的goroutine其中一个,不要求持有lock
	// 唤醒时goroutine可能不在wait状态,这时可能导致某些问题
	// 可能需要别的检测状态或设置一个timer持续唤醒,也可用for套住Signal
	time.Sleep(time.Second)
	cond.Signal()
	time.Sleep(time.Second)
	// 唤醒所有在wait的goroutine
	cond.Broadcast()
	time.Sleep(time.Second)
}
