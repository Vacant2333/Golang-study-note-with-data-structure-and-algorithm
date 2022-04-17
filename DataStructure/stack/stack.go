package stack

import (
	"errors"
)

/*
	stack 栈 用链表实现 头结点不存数据 存长度
*/

type Node struct {
	Data int
	Next *Node
}

// Create 创建一个栈 返回head
func Create() *Node {
	head := new(Node)
	head.Next = nil
	head.Data = 0
	return head
}

// Push 向一个栈中插入Node
func (stack *Node) Push(data int) {
	// 计数++
	stack.Data++
	cur := stack
	// 新Node
	tmp := new(Node)
	tmp.Data = data
	tmp.Next = nil
	// 移动到最后一个Node
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = tmp
}

// Pop 从一个栈中拿出数据
func (stack *Node) Pop(delete bool) (int, error) {
	if stack.isEmpty() == false {
		cur := stack
		// 栈不为空,移动到倒数第二个Node
		for cur.Next.Next != nil {
			cur = cur.Next
		}
		// 拿到值然后断开Next
		tmp := cur.Next.Data
		// delete是true的时候删除Node
		if delete == true {
			stack.Data--
			cur.Next = nil
		}
		return tmp, nil
	} else {
		// 栈为空
		return 0, errors.New("stack is empty")
	}
}

// isEmpty 检查栈是否为空
func (stack *Node) isEmpty() bool {
	return stack.Data == 0
}
