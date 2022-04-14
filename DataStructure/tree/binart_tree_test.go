package tree

import (
	"fmt"
	"testing"
)

func TestCreateNode(t *testing.T) {
	n := CreateNode("10", nil, nil)
	fmt.Println(n)
}

func TestNode_CountSon(t *testing.T) {
	root := CreateNode("0", nil, nil)
	root.Right = CreateNode("1", nil, nil)
	fmt.Println(root.CountSon())
}

func TestNode_PrintTree(t *testing.T) {
	root := CreateNode("0", nil, nil)
	root.Right = CreateNode("1", nil, nil)
	root.PrintTree()
}

func TestNode_PrintNode(t *testing.T) {
	root := CreateNode("0", nil, nil)
	root.PrintNode()
}

func TestNode_GetDepth(t *testing.T) {
	root := CreateNode("0", nil, nil)
	root.Right = CreateNode("1", nil, nil)
	fmt.Println(root.GetDepth(), root.Right.GetDepth())
}
