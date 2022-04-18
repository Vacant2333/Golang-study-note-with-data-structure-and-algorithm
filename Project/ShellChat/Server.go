package main

import (
	"errors"
	"fmt"
	"net"
)

type Client struct {
	name string
	conn net.Conn
	addr string
}

var (
	clients = make([]Client, 0)
)

func main() {
	// 监听8000端口
	listener, e := net.Listen("tcp", "127.0.0.1:8000")
	if e != nil {
		panic(errors.New("net listen error"))
	}
	// 接受用户登入
	for {
		conn, e := listener.Accept()
		if e != nil {
			panic(errors.New("listener accept error"))
		}
		// 保存用户的连接
		client := Client{"test", conn, conn.RemoteAddr().String()}
		clients = append(clients, client)
		fmt.Println("User login", client)
		// 每一个用户的连接在单独的协程中处理
		go listenClient(client)
	}
}

// 与一个client通信
func listenClient(client Client) {
	buffer := make([]byte, 256)
	count, _ := client.conn.Read(buffer)
	if count > 0 {
		fmt.Println(count, buffer)
	}
}
