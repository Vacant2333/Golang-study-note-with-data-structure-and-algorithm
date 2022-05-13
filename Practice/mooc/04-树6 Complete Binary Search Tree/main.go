package main

import (
	"fmt"
	"math"
	"sort"
)

/*
A Binary Search Tree (BST) is recursively defined as a binary tree which has the following properties:
The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
Both the left and right subtrees must also be binary search trees.
A Complete Binary Tree (CBT) is a tree that is completely filled, with the possible exception of
the bottom level, which is filled from left to right.
Now given a sequence of distinct non-negative integer keys, a unique BST can be constructed
if it is required that the tree must also be a CBT. You are supposed to output the level order
traversal sequence of this BST.
Input Specification:
Each input file contains one test case. For each case, the first line contains a positive integer
N (≤1000). Then N distinct non-negative integer keys are given in the next line. All the numbers
in a line are separated by a space and are no greater than 2000.
Output Specification:
For each test case, print in one line the level order traversal sequence of the corresponding
complete binary search tree. All the numbers in a line must be separated by a space, and there
must be no extra space at the end of the line.
Sample Input:
10
1 2 3 4 5 6 7 8 9 0
Sample Output:
6 3 8 1 5 7 9 0 2 4
*/

func main() {
	// 读入数据
	var count, tmp int
	fmt.Scan(&count)
	sourceTree := make([]int, count)
	resultTree := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&tmp)
		sourceTree[i] = tmp
	}
	sort.Ints(sourceTree)
	// 主要的处理函数
	var solve func(left, right, root int)
	solve = func(left, right, root int) {
		// 这一段的元素总个数
		n := right - left + 1
		// 没有元素可以处理 直接return
		if n <= 0 {
			return
		}
		// 左子树的规模(有多少个节点)
		l := getLeftLength(n)
		resultTree[root] = sourceTree[left+l]
		leftTRoot := root*2 + 1
		rightTRoot := leftTRoot + 1
		solve(left, left+l-1, leftTRoot)
		solve(left+l+1, right, rightTRoot)
	}
	solve(0, len(sourceTree)-1, 0)
	for i := 0; i < count; i++ {
		fmt.Print(resultTree[i])
		if i != count-1 {
			fmt.Print(" ")
		}
	}
}

// 计算n个节点的树左子树有多少个节点
func getLeftLength(n int) int {
	// 树的完全部分的高度
	h := math.Floor(math.Log2(float64(n + 1)))
	// 最下面一层的节点数量
	x := float64(n) - (math.Pow(2, h) - 1)
	// x最大只能有最下面一层的一半
	x = math.Min(x, math.Pow(2, h-1))
	// 左子树的节点数量
	l := math.Pow(2, h-1) - 1 + x
	return int(l)
}
