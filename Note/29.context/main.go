package main

import (
	"context"
	"fmt"
)

/*
	https://golang.design/go-questions/stdlib/context/why/
	context 译作上下文,准确说是goroutine的上下文
	主要用来在goroutine之间传递上下文信息,包括:取消信号,超时时间,截止时间,kv数据

*/

func main() {
	ctx := context.Background()
	// 模拟建立一个request
	ctx = context.WithValue(ctx, "requestId", "request12345")
	handleRequest(ctx)
}

func handleRequest(ctx context.Context) {
	requestId, ok := ctx.Value("requestId").(string)
	if ok {
		fmt.Printf("process request: %v\n", requestId)
	}
}
