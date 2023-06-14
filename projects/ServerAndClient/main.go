package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Server struct {
	name string
	ch   chan int
}

type Client struct {
	name string
	ch   chan int
}

func MakeServer(cname string, ch chan int) {
	//请在此处实现消费者的代码
	server := &Server{
		name: cname,
		ch:   ch,
	}
	// 处理请求
	go server.serve()
}

func (s *Server) serve() {
	for v := range s.ch {
		fmt.Printf("S[%v] get a value[%v]\n", s.name, v)
	}
}

func MakeClient(pname string, ch chan int) {
	//请在此处实现生产者的代码
	client := &Client{
		name: pname,
		ch:   ch,
	}
	go client.sendRandNums()
}

func (c *Client) sendRandNums() {
	max := 10000
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		value := rand.Int() % max
		go func() {
			fmt.Printf("C[%v] put a value[%v]\n", c.name, value)
			c.ch <- value
			wg.Done()
		}()
	}
	wg.Wait()
}

func main() {
	//用channel来传递"产品", 不再需要自己去加锁维护一个全局的阻塞队列
	ch := make(chan int)
	go func() {
		for i := 0; i < 5000; i++ {
			go MakeClient(fmt.Sprintf("生产者[%v]", i), ch)
		}
		for i := 0; i < 10; i++ {
			go MakeServer(fmt.Sprintf("消费者[%v]", i), ch)
		}
	}()

	time.Sleep(20 * time.Second)
	close(ch)
	time.Sleep(10 * time.Second)
}
