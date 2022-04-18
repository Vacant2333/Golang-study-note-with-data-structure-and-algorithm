package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
)

func main() {
	// 连接服务端
	conn, e := net.Dial("tcp", "127.0.0.1:8000")
	if e != nil {
		panic(errors.New("tcp connect error"))
	}
	// 退出时关闭连接
	defer func() {
		conn.Close()
	}()
	// 处理用户输入
	go handleSend(conn)
	// 处理服务器输出
	go handleReceive(conn)
	// 挂起主进程
	for {
	}
}

// 协程处理用户发送信息
func handleSend(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, _ := reader.ReadLine()
		_, err := conn.Write(line)
		if err != nil {
			panic(errors.New("conn write error"))
		}
	}
}

// 协程处理服务端信息
func handleReceive(conn net.Conn) {
	buffer := make([]byte, 256)
	for {
		conn.Read(buffer)
		fmt.Println(string(buffer))
	}
}
