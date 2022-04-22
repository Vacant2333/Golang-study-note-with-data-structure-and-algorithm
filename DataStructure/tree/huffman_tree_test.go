package tree

import (
	"testing"
)

func TestBuildHuffmanTree(t *testing.T) {
	tree := BuildHuffmanTree([]ElementType{"1", "2", "3", "4", "5"})
	tree.PrintTree()
}
