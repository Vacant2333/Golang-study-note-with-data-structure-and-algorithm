package main

import (
	"bufio"
	"errors"
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
	// 单独的协程处理用户输入
	go handleSend(conn)

	for {

	}
}

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
