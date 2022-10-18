package main

import "fmt"

/*
	闭包就是 一个函数 和 与其相关的引用环境 组合的一个整体(实体)
	addUpper 是一个函数 返回的数据类型是 func(int) int
*/
/*
func addUpper() func(int) int {
	var n = 10
	return func(x int) int {
		n = n + x
		return n
	}
}

func main() {
	// f变量的内容其实就是addUpper中 return的内容 也就是那个匿名函数
	f := addUpper()
	// 返回11 也就是 n(10) + x(1)
	fmt.Println(f(1))
	// 返回13 也就是 n(11) + x(2)
	// n在上一次运行的时候已经被赋值为11 但是这个时候不会释放addUpper,所以n也不会被释放
	// 也可以理解成匿名函数和n形成了一个整体,构成了闭包
	fmt.Println(f(2))
	// 返回16 n在上一次已经为13了
	fmt.Println(f(3))
}
*/
type Node struct {
	Val  int
	Next *Node
}

func main() {
	a1 := []int{1, 2}
	l1 := buildList(a1)
	r(l1)
	printList(l1)
}

func buildList(v []int) *Node {
	head := new(Node)
	head.Val = v[0]
	cur := head
	for i := 1; i < len(v); i++ {
		tmp := new(Node)
		tmp.Val = v[i]
		cur.Next = tmp
		cur = tmp
	}
	cur.Next = head
	return head
}

func printList(head *Node) {
	cur := head.Next
	fmt.Print(head.Val, " ")
	for cur != head {
		fmt.Print(cur.Val, " ")
		cur = cur.Next
	}
	fmt.Print(head.Val, " ")
}

func r(head *Node) {
	// 如果只有一个节点或者没有节点
	if head == nil || head.Next == head {
		return
	}
	// 1
	pre := head
	// 2
	cur := head.Next
	// count为1/0 为1时删除当前的Node
	count := 1
	for cur != head {
		if count == 1 {
			// 删除当前节点（cur）
			count = 0
			tmp := cur.Next
			pre.Next = tmp
			cur = tmp
		} else {
			// count为0 不删除节点，移动cur
			cur = cur.Next
			pre = pre.Next
			count = 1
		}
	}
}
