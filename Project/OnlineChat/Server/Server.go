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
	clients   = make([]Client, 0)
	msgRecord = make([]string, 0)
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
		name := make([]byte, 128)
		conn.Read(name)
		client := Client{string(name), conn, conn.RemoteAddr().String()}
		// 保存客户端
		clients = append(clients, client)
		fmt.Println("User login:", client)
		// 每一个用户的连接在单独的协程中处理
		go listenClient(client)
	}
}

// 与一个client通信
func listenClient(client Client) {
	syncUserRecord(client)
	for {
		buffer := make([]byte, 128)
		count, e := client.conn.Read(buffer)
		if e != nil && e.Error() == "EOF" {
			// 连接已关闭
			fmt.Println("User logout:", client)
			closeClientConn(client)
			return
		}
		fmt.Println(e)
		if count > 0 {
			saveMsg(formatMsg(client.name, buffer))
			sendMsgToAll(client.name, buffer)
		}
	}
}

// 发送信息给所有用户
func sendMsgToAll(name string, msg []byte) {
	m := formatMsg(name, msg)
	for _, client := range clients {
		client.conn.Write([]byte(m))
	}
}

// 同步用户消息记录
func syncUserRecord(client Client) {
	for _, msg := range msgRecord {
		client.conn.Write([]byte(msg))
	}
}

// 获得消息格式
func formatMsg(name string, msg []byte) string {
	return name + ": " + string(msg) + "\n"
}

// 保存消息记录
func saveMsg(msg string) {
	msgRecord = append(msgRecord, msg)
}

// 关闭连接
func closeClientConn(client Client) {
	client.conn.Close()
	for i, v := range clients {
		if v == client {
			clients = append(clients[:i], clients[i+1:]...)
			break
		}
	}
}
