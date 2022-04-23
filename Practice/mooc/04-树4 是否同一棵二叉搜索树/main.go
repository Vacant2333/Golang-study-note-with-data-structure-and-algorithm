package main

import (
	"fmt"
)

/*
https://pintia.cn/problem-sets/1497448825169559552/problems/1505905647748263937
给定一个插入序列就可以唯一确定一棵二叉搜索树。然而，一棵给定的二叉搜索树却可以由多种不同的插入序列得到。
例如分别按照序列{2, 1, 3}和{2, 3, 1}插入初始为空的二叉搜索树，都得到一样的结果。于是对于输入的各种插入序列，
你需要判断它们是否能生成一样的二叉搜索树。
输入格式:
输入包含若干组测试数据。每组数据的第1行给出两个正整数N (≤10)和L，分别是每个序列插入元素的个数和需要检查的序列个数。
第2行给出N个以空格分隔的正整数，作为初始插入序列。随后L行，每行给出N个插入的元素，属于L个需要检查的序列。
简单起见，我们保证每个插入序列都是1到N的一个排列。当读到N为0时，标志输入结束，这组数据不要处理。
输出格式:
对每一组需要检查的序列，如果其生成的二叉搜索树跟对应的初始序列生成的一样，输出“Yes”，否则输出“No”。
输入样例:
4 2
3 1 4 2
3 4 1 2
3 2 4 1
2 1
2 1
1 2
0
输出样例:
Yes
No
No
*/
func main() {
	for {
		var nodeCount, checkCount int
		var tmp ElementType
		// 节点数量
		fmt.Scan(&nodeCount)
		if nodeCount == 0 {
			return
		}
		// 要比较的搜索树数量
		fmt.Scan(&checkCount)
		// 读入树（总数是要比较的搜索树的数量+1）
		// 第一棵树 用来和后面的树比较
		var firstTree *Node
		for a := 0; a <= checkCount; a++ {
			// 读入的节点数据
			var nodeTmp []ElementType
			var treeTmp *Node
			for i := 0; i < nodeCount; i++ {
				// 读入节点
				fmt.Scan(&tmp)
				nodeTmp = append(nodeTmp, tmp)
			}
			if a == 0 {
				firstTree = BuildBSTFromSlice(nodeTmp)
			} else {
				// 和第一个树比较
				treeTmp = BuildBSTFromSlice(nodeTmp)
				if IsEqual(firstTree, treeTmp) == false {
					// 和第一个树不一样
					fmt.Println("No")
				} else {
					fmt.Println("Yes")
				}
			}
		}
	}
}

// IsEqual 比较两个树是否完全一样
func IsEqual(root1 *Node, root2 *Node) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}
	return root1.Data == root2.Data && IsEqual(root1.Left, root2.Left) && IsEqual(root1.Right, root2.Right)
}

// BuildBSTFromSlice 从切片生成一个搜索二叉树BST
func BuildBSTFromSlice(s []ElementType) *Node {
	// 根节点
	t := CreateNode(s[0], nil, nil)
	// 插入其余元素
	for i := 1; i < len(s); i++ {
		t.InsertBSTNode(s[i])
	}
	return t
}

type ElementType int

type Node struct {
	Data  ElementType
	Left  *Node
	Right *Node
}

// CreateNode 创建一个二叉树节点/根
func CreateNode(data ElementType, left, right *Node) *Node {
	tree := new(Node)
	tree.Data = data
	tree.Left = left
	tree.Right = right
	return tree
}

func (root *Node) InsertBSTNode(data ElementType) {
	cur := root
	for cur.Data != data {
		if data < cur.Data {
			// 要插入的节点小于当前节点
			if cur.Left != nil {
				// 左节点不为空 移动cur
				cur = cur.Left
			} else {
				// 左节点是空的 插入新节点后退出
				cur.Left = CreateNode(data, nil, nil)
				return
			}
		} else {
			// 要插入的节点大于当前节点
			if cur.Right != nil {
				// 右节点不为空 移动cur
				cur = cur.Right
			} else {
				// 左节点是空的 插入新节点后退出
				cur.Right = CreateNode(data, nil, nil)
				return
			}
		}
	}
	// 到了这里就是t和data相等 不作任何操作 直接退出
}
