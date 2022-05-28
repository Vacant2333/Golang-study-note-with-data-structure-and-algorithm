package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	var ip string
	fmt.Println("Input Address:")
	fmt.Scan(&ip)
	wg := &sync.WaitGroup{}
	wg.Add(65535)
	for i := 1; i <= 65535; i++ {
		go checkPort(ip, i, wg)
	}
	wg.Wait()
}

func checkPort(ip string, port int, wg *sync.WaitGroup) {
	add := fmt.Sprintf("%v:%v", ip, port)
	_, err1 := net.DialTimeout("tcp", add, time.Second*15)
	if err1 == nil {
		fmt.Printf("%v[TCP] Open!\n", add)
	}
	wg.Done()
}
