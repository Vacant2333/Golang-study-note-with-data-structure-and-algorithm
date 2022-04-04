package tree

import "fmt"

/*
	二叉树,每一个节点可以认为是一个二叉树
	实现:
		二叉搜索树 待完成
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
func (node *Node) CountSon() int {
	count := 0
	if node.Left != nil {
		count++
	}
	if node.Right != nil {
		count++
	}
	return count
}

// PrintTree 打印整棵树(前序)
func (node *Node) PrintTree() {
	if node != nil {
		node.PrintNode()
		node.Left.PrintTree()
		node.Right.PrintTree()
	}
}

// PrintNode 打印树节点
func (node *Node) PrintNode() {
	if node != nil {
		if node.Left != nil && node.Right != nil {
			// 左右都不为空
			fmt.Printf("Node:%v Left:%v: Right:%v\n", node.Data, node.Left.Data, node.Right.Data)
		} else if node.Left != nil && node.Right == nil {
			// 只有左节点
			fmt.Printf("Node:%v Left:%v\n", node.Data, node.Left.Data)
		} else if node.Left == nil && node.Right != nil {
			// 只有右节点
			fmt.Printf("Node:%v Right:%v\n", node.Data, node.Right.Data)
		}
	}
}