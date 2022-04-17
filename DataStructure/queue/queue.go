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
func (queue *Node) Push(data int) {
	// 计数+1
	queue.Data++
	cur := queue
	// 声明新节点
	tmp := new(Node)
	tmp.Data = data
	tmp.Next = nil
	// 移动cur到最后一个节点
	for cur.Next != nil {
		cur = cur.Next
	}
	// 插入新节点
	cur.Next = tmp
}

// Pop 从队列头部拿出一个节点(delete为true时删除)
func (queue *Node) Pop(delete bool) (int, error) {
	if queue.isEmpty() == false {
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
func (queue *Node) isEmpty() bool {
	return queue.Data == 0
}
