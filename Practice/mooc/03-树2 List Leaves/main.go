package main

import (
	"fmt"
	"strconv"
)

/*
https://pintia.cn/problem-sets/1497448825169559552/problems/1503334324290916353
Given a tree, you are supposed to list all the leaves in the order of top down, and left to right.
Input Specification:
Each input file contains one test case. For each case,
the first line gives a positive integer N (≤10) which is the total number of nodes in the tree -- and
hence the nodes are numbered from 0 to N−1. Then N lines follow, each corresponds to a node, and gives
the indices of the left and right children of the node. If the child does not exist, a "-" will be put
at the position. Any pair of children are separated by a space.
Output Specification:
For each test case, print in one line all the leaves' indices in the order of top down, and left to right.
There must be exactly one space between any adjacent numbers, and no extra space at the end of the line.
Sample Input:
8
1 -
- -
0 -
2 7
- -
- -
5 -
4 6
Sample Output:
4 1 5
*/

type Node struct {
	data  int
	left  *Node
	right *Node
}

type NodeData struct {
	left  int
	right int
}

func main() {
	var nodeCount int
	fmt.Scan(&nodeCount)
	// 题目不提供head节点，在所有数据中没出现的下标就是head
	headMap := make(map[int]bool, 0)
	for i := 0; i < nodeCount; i++ {
		// 全部设为true 代表可能是head
		headMap[i] = true
	}
	dataSlice := make([]NodeData, 0)
	var left, right string
	var l, r int
	for i := 0; i < nodeCount; i++ {
		// 读入左右节点 如果是-的话就为空
		fmt.Scan(&left, &right)
		tmp := NodeData{}
		if left != "-" {
			l, _ = strconv.Atoi(left)
			tmp.left = l
			headMap[l] = false
		} else {
			tmp.left = -1
		}
		if right != "-" {
			r, _ = strconv.Atoi(right)
			tmp.right = r
			headMap[r] = false
		} else {
			tmp.right = -1
		}
		dataSlice = append(dataSlice, tmp)
	}
	// 找头节点！
	headIndex := 0
	for i := 0; i < nodeCount; i++ {
		if headMap[i] == true {
			headIndex = i
			break
		}
	}
	// 建树
	head := build(headIndex, dataSlice)
	// 输出叶节点
	printTree(head)
}

// 建树 返回head
func build(headIndex int, s []NodeData) *Node {
	head := new(Node)
	head.data = headIndex
	if s[headIndex].left != -1 {
		head.left = build(s[headIndex].left, s)
	}
	if s[headIndex].right != -1 {
		head.right = build(s[headIndex].right, s)
	}
	return head
}

// 从上到下从左到右打印叶节点
func printTree(head *Node) {
	// 用来记录当前层在nodes中的下标
	left := 0
	nodes := make([]*Node, 0)
	nodes = append(nodes, head)
	out := make([]int, 0)
	for {
		// 先检查当前层有没有叶节点 有的话输出
		for i := left; i < len(nodes); i++ {
			if nodes[i].left == nil && nodes[i].right == nil {
				out = append(out, nodes[i].data)
			}
		}
		tmp := len(nodes)
		// 把当前层的节点加入到nodes中
		for i := left; i < tmp; i++ {
			if nodes[i].left != nil {
				nodes = append(nodes, nodes[i].left)
			}
			if nodes[i].right != nil {
				nodes = append(nodes, nodes[i].right)
			}
		}
		// 如果这次没加入新节点就退出
		if len(nodes) == left {
			break
		}
		// 修改left
		left = tmp
	}
	// 输出out
	for i := 0; i < len(out); i++ {
		fmt.Print(out[i])
		if i != len(out)-1 {
			fmt.Print(" ")
		}
	}
}
