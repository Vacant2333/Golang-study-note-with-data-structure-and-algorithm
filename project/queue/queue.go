package queue

import "errors"

type Node struct {
	Data int
	Next *Node
}

/*
队列,先入先出,插入从尾插,拿从头拿
head节点存长度不存数据
*/

// Create 创建一个队列 返回头结点
func Create() *Node {
	queue := new(Node)
	queue.Data = 0
	queue.Next = nil
	return queue
}

// Push 向队列尾部插入一个节点
func Push(queue *Node, data int) {
	queue.Data++
	tmp := new(Node)
	tmp.Data = data
	tmp.Next = nil
	for queue.Next != nil {
		queue = queue.Next
	}
	queue.Next = tmp
}

// Pop 从队列头部拿出一个节点(delete为true时删除)
func Pop(queue *Node, delete bool) (int, error) {
	if isEmpty(queue) == false {
		tmp := queue.Next
		if delete == true {
			queue.Next = queue.Next.Next
			queue.Data--
		}
		return tmp.Data, nil
	} else {
		return 0, errors.New("queue is empty")
	}
}

// isEmpty 检查队列是否为空
func isEmpty(queue *Node) bool {
	return queue.Data == 0
}
