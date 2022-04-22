package tree

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

var s = []ElementType{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "July", "Aug", "Sep", "Oct", "Nov", "Dec"}

func TestCreateBSTFromSlice(t *testing.T) {
	var data []ElementType
	count := 1000 * 10000
	// 存入随机数
	for i := 0; i < count; i++ {
		data = append(data, ElementType(strconv.Itoa(rand.Intn(10000000000))))
	}
	// 生成BST
	bst := BuildBSTFromSlice(data)
	// 查找测试
	for i := 0; i < count; i++ {
		if bst.FindBSTNode(data[i]) == nil {
			fmt.Println("Find Error")
		}
	}
	// BST验证
	fmt.Println("Valid:", bst.IsValidBST())
	fmt.Println("Depth:", bst.GetDepth())
}

func TestNode_FindBSTNode(t *testing.T) {
	bst := BuildBSTFromSlice(s)
	for _, v := range s {
		fmt.Println(bst.FindBSTNode(v))
	}
}

func TestNode_FindMinBSTNode(t *testing.T) {
	bst := BuildBSTFromSlice(s)
	fmt.Println(bst.FindMinBSTNode())
}

func TestNode_FindMaxBSTNode(t *testing.T) {
	bst := BuildBSTFromSlice(s)
	fmt.Println(bst.FindMaxBSTNode())
}

func TestNode_InsertBSTNode(t *testing.T) {
	bst := BuildBSTFromSlice(s)
	bst.InsertBSTNode("test")
	bst.InsertBSTNode("china")
	bst.InsertBSTNode("牛")
	bst.PrintTree()
	fmt.Println(bst.IsValidBST())
}

func TestNode_DeleteBSTNode(t *testing.T) {
	bst := BuildBSTFromSlice(s)
	// 没有儿子节点
	bst.DeleteBSTNode("Dec")
	bst.PrintTree()
	fmt.Printf("valid:%v\n", bst.IsValidBST())

	bst = BuildBSTFromSlice(s)
	// 一个儿子节点
	bst.DeleteBSTNode("Aug")
	bst.PrintTree()
	fmt.Printf("valid:%v\n", bst.IsValidBST())

	bst = BuildBSTFromSlice(s)
	// 两个儿子节点
	bst.DeleteBSTNode("Mar")
	bst.PrintTree()
	fmt.Printf("valid:%v\n", bst.IsValidBST())
}

func TestNode_IsValidBST(t *testing.T) {
	bst := BuildBSTFromSlice(s)
	fmt.Println(bst.IsValidBST())
	bst.Left.Data = "0"
	fmt.Println(bst.IsValidBST())
}
