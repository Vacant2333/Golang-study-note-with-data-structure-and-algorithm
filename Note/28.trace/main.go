package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, _ := os.Create("trace.out")
	trace.Start(f)
	defer trace.Stop()

	fmt.Println("trace print")
	// 使用go tool trace trace.out 命令
}
