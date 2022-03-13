package Stack

import "errors"

/*
	Stack 栈 用链表实现 头结点不存数据 存长度
*/

type ElementType = int32

type Node struct {
	Data ElementType
	Next *Node
}

// Create 创建一个栈 返回head
func Create() *Node {
	head := new(Node)
	head.Next = nil
	head.Data = 0
	return head
}

// Push 向一个栈中插入数据
func Push(stack *Node, data ElementType) {
	stack.Data++
	tmp := new(Node)
	tmp.Data = data
	tmp.Next = nil
	// 移动到最后一个Node
	for stack.Next != nil {
		stack = stack.Next
	}
	stack.Next = tmp
}

// Pop 从一个栈中拿出数据
func Pop(stack *Node) (ElementType, error) {
	if isNull(stack) == false {
		stack.Data--
		cur := stack
		// 栈不为空,移动到倒数第二个Node
		for cur.Next.Next != nil {
			cur = cur.Next
		}
		// 拿到值然后断开Next
		tmp := cur.Next.Data
		cur.Next = nil
		return tmp, nil
	} else {
		// 栈为空
		return 0, errors.New("stack is empty")
	}
}

// isNull 检查栈是否为空
func isNull(stack *Node) bool {
	return stack.Data == 0
}
