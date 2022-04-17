package queue

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	head := Create()
	fmt.Println(head)
}

func TestNode_Push(t *testing.T) {
	head := Create()
	head.Push(1)
	head.Push(2)
	head.Push(3)
	fmt.Println(head.Next)
}

func TestNode_Pop(t *testing.T) {
	head := Create()
	head.Push(1)
	head.Push(2)
	head.Push(3)
	fmt.Println(head.Pop(true))
	fmt.Println(head.Pop(true))
	fmt.Println(head.Pop(true))
}
