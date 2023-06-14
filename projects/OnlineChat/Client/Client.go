package main

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"
)

var stop chan bool

func main() {
	// 登陆用户名
	var name string
	fmt.Println("请输入用户名:")
	fmt.Scanf("%s", &name)
	// 连接服务端
	conn, e := net.Dial("tcp", "127.0.0.1:8000")
	if e != nil {
		panic(errors.New("tcp connect error"))
	}
	fmt.Println("连接服务器成功")
	// 退出时关闭连接
	defer func() {
		conn.Close()
	}()
	// 处理用户输入
	go handleSend(conn, name)
	// 处理服务器输出
	go handleReceive(conn)
	// 挂起主进程
	<-stop
}

// 协程处理用户发送信息
func handleSend(conn net.Conn, name string) {
	// 设置用户名
	conn.Write([]byte(name))
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
	for {
		buffer := make([]byte, 128)
		n, _ := conn.Read(buffer)
		if n > 0 {
			fmt.Printf("%s", buffer)
		}
	}
}
