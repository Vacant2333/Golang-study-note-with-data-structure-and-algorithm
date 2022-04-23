package main

import (
	"fmt"
)

/*
https://pintia.cn/problem-sets/1497448825169559552/problems/1503334324290916354
An inorder binary tree traversal can be implemented in a non-recursive way with a stack.
For example, suppose that when a 6-node binary tree (with the keys numbered from 1 to 6) is traversed,
the stack operations are: push(1); push(2); push(3); pop(); pop(); push(4); pop(); pop(); push(5); push(6); pop(); pop().
Then a unique binary tree (shown in Figure 1) can be generated from this sequence of operations.
Your task is to give the postorder traversal sequence of this tree.
Figure 1
Input Specification:
Each input file contains one test case. For each case, the first line contains a positive integer N (≤30)
which is the total number of nodes in a tree (and hence the nodes are numbered from 1 to N).
Then 2N lines follow, each describes a stack operation in the format: "Push X" where X is
the index of the node being pushed onto the stack; or "Pop" meaning to pop one node from the stack.
Output Specification:
For each test case, print the postorder traversal sequence of the corresponding tree in one line.
A solution is guaranteed to exist. All the numbers must be separated by exactly one space,
and there must be no extra space at the end of the line.
Sample Input:
6
Push 1
Push 2
Push 3
Pop
Pop
Push 4
Pop
Pop
Push 5
Push 6
Pop
Pop
Sample Output:
3 4 2 6 5 1
*/
var nodeCount int
var nodeMax int
var printCount int

func main() {
	fmt.Scan(&nodeMax)
	nodeCount = 0
	printCount = 0
	t := buildTreeFromInput()
	printTree(t)
}

// 后序打印树
func printTree(node *Node) {
	if node != nil {
		printTree(node.left)
		printTree(node.right)
		if printCount == nodeMax-1 {
			fmt.Print(node.data)
		} else {
			printCount++
			fmt.Print(node.data, " ")
		}
	}
}

func buildTreeFromInput() *Node {
	if nodeCount == nodeMax {
		return nil
	}
	data, t := parseInput()
	if t == true {
		// Push操作,data有数据
		nodeCount++
		node := new(Node)
		node.data = data
		node.left = buildTreeFromInput()
		node.right = buildTreeFromInput()
		return node
	}
	return nil
}

// 解析输入 true是push,false是pop
func parseInput() (int, bool) {
	var command string
	var data int
	fmt.Scan(&command)
	if command == "Push" {
		fmt.Scan(&data)
		return data, true
	}
	return -1, false
}

type Node struct {
	data        int
	left, right *Node
}
