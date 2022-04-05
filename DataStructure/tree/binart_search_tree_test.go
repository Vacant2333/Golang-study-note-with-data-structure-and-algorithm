package tree

import (
	"fmt"
	"testing"
)

var s = []ElementType{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "July", "Aug", "Sep", "Oct", "Nov", "Dec"}

func TestCreateBSTFromSlice(t *testing.T) {
	bst := CreateBSTFromSlice(s)
	bst.PrintTree()

	//var a []ElementType TODO:speed test
}

func TestNode_FindBSTNode(t *testing.T) {
	bst := CreateBSTFromSlice(s)
	for _, v := range s {
		fmt.Println(bst.FindBSTNode(v))
	}
}

func TestNode_FindMinBSTNode(t *testing.T) {
	bst := CreateBSTFromSlice(s)
	fmt.Println(bst.FindMinBSTNode())
}

func TestNode_FindMaxBSTNode(t *testing.T) {
	bst := CreateBSTFromSlice(s)
	fmt.Println(bst.FindMaxBSTNode())
}

func TestNode_InsertBSTNode(t *testing.T) {
	// TODO:check
}

func TestNode_IsValidBST(t *testing.T) {
	bst := CreateBSTFromSlice(s)
	fmt.Println(bst.IsValidBST())
	bst.Left.Data = "0"
	fmt.Println(bst.IsValidBST())
}
