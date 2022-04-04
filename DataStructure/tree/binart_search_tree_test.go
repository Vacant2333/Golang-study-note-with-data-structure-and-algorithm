package tree

import (
	"fmt"
	"testing"
)

var s = []ElementType{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "July", "Aug", "Sep", "Oct", "Nov", "Dec"}

func TestCreateBSTFromSlice(t *testing.T) {
	bst := CreateBSTFromSlice(s)
	bst.PrintTree()
}

func TestFindBSTNode(t *testing.T) {
	bst := CreateBSTFromSlice(s)
	for _, v := range s {
		fmt.Println(bst.FindBSTNode(v))
	}
}
