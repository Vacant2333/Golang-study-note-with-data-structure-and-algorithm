package main

import (
	"Golang-study-note-with-data-structure-and-algorithm/DataStructure/queue"
	"fmt"
)

func main() {
	q := queue.Create()
	queue.Push(q, 1)
	queue.Push(q, 2)
	queue.Push(q, 3)
	fmt.Println(queue.Pop(q, true))
	fmt.Println(queue.Pop(q, true))
	queue.Push(q, 4)
	fmt.Println(queue.Pop(q, true))
	fmt.Println(queue.Pop(q, true))
	fmt.Println(q.Data)
}
