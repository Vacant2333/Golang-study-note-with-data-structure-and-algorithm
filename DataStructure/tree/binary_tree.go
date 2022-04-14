package tree

import "fmt"

/*
	公式:
		(度: 某节点的儿子的数量 度0就是叶子节点)
		总结点            = 总的度(总线) + 1
		度0              = 度2 + 1
		第i层最大节点数    = 2^(i-1)
		深度为k的最大节点数 = 2^k - 1
		满二叉树(除了叶结点外每一个结点都有左右子叶且叶结点都处在最底层的二叉树):
			第k层的节点数    = 2^(k-1)
			深度为m的总结点数 = 2^m - 1
			总结点为n的深度   = log2(n + 1)

	二叉树,每一个节点可以认为是一个二叉树
	实现:
		二叉搜索树(Binary Search Tree) 已完成
		平衡二叉搜索树(AVL)             TODO: AVL
		红黑树(Red Black Tree)        TODO: RBT
*/

type ElementType string

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

// CountSon 计算子节点数量(0/1/2)
func (root *Node) CountSon() int {
	count := 0
	if root.Left != nil {
		count++
	}
	if root.Right != nil {
		count++
	}
	return count
}

// PrintTree 打印整棵树(前序)
func (root *Node) PrintTree() {
	if root != nil {
		root.PrintNode()
		root.Left.PrintTree()
		root.Right.PrintTree()
	}
}

// PrintNode 打印树节点
func (root *Node) PrintNode() {
	if root != nil {
		if root.Left != nil && root.Right != nil {
			// 左右都不为空
			fmt.Printf("(2)Node:%v Left:%v: Right:%v\n", root.Data, root.Left.Data, root.Right.Data)
		} else if root.Left != nil && root.Right == nil {
			// 只有左节点
			fmt.Printf("(1)Node:%v Left:%v\n", root.Data, root.Left.Data)
		} else if root.Left == nil && root.Right != nil {
			// 只有右节点
			fmt.Printf("(1)Node:%v Right:%v\n", root.Data, root.Right.Data)
		} else {
			// 没有子节点
			fmt.Printf("(0)Node:%v\n", root.Data)
		}
	}
}

// GetDepth 获得树的深度
func (root *Node) GetDepth() int {
	if root == nil {
		return 0
	} else {
		leftMaxDepth := root.Left.GetDepth()
		rightMaxDepth := root.Right.GetDepth()
		if leftMaxDepth >= rightMaxDepth {
			return leftMaxDepth + 1
		} else {
			return rightMaxDepth + 1
		}
	}
}
