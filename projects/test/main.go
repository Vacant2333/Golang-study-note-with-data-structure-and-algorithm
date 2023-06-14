package main

import (
	"fmt"
	_ "fmt"
)

type Node struct {
	value int
	next  *Node
}

func main() {
	fmt.Println()
}

func addNode(head *Node, value int) {
	head.next = &Node{
		value: value,
		next:  nil,
	}
}
