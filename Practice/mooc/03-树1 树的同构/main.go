package main

import (
	"Golang-study-note-with-data-structure-and-algorithm/DataStructure/tree"
	"fmt"
	"strconv"
)

/*
	url : https://pintia.cn/problem-sets/1497448825169559552/problems/1503334324290916352
	给定两棵树T1和T2。如果T1可以通过若干次左右 孩子互换就变成T2，则我们称两棵树是“同构”的。例如图1给出的两棵树就是同构的，因为我们把其中一棵树的结点A、B、G的左右 孩子互换后，就得到另外一棵树。而图2就不是同构的。
8
A 1 2
B 3 4
C 5 -
D - -
E 6 -
G 7 -
F - -
H - -
8
G - 4
B 7 6
F - -
A 5 1
H - -
C 0 -
D - -
E 2 -
*/

type nodeData struct {
	data       tree.ElementType
	leftIndex  int
	rightIndex int
}

func main() {
	var tree1Count, tree2Count int
	var tree1AIndex, tree2AIndex int
	var tmp string
	fmt.Scan(&tree1Count)
	// 将所有节点的信息存入slice
	tree1Data := make([]nodeData, tree1Count)
	for i := 0; i < tree1Count; i++ {
		tree1Data[i] = *new(nodeData)
		fmt.Scan(&tree1Data[i].data)
		if tree1Data[i].data == "A" {
			tree1AIndex = i
		}
		fmt.Scan(&tmp)
		if tmp != "-" {
			tree1Data[i].leftIndex, _ = strconv.Atoi(tmp)
		} else {
			tree1Data[i].leftIndex = -1
		}
		fmt.Scan(&tmp)
		if tmp != "-" {
			tree1Data[i].rightIndex, _ = strconv.Atoi(tmp)
		} else {
			tree1Data[i].rightIndex = -1
		}
	}
	tree1 := generateTree(tree1Data, tree1AIndex)

	fmt.Scan(&tree2Count)
	// 将所有节点的信息存入slice
	tree2Data := make([]nodeData, tree2Count)
	for i := 0; i < tree2Count; i++ {
		tree2Data[i] = *new(nodeData)
		fmt.Scan(&tree2Data[i].data)
		if tree2Data[i].data == "A" {
			tree2AIndex = i
		}
		fmt.Scan(&tmp)
		if tmp != "-" {
			tree2Data[i].leftIndex, _ = strconv.Atoi(tmp)
		} else {
			tree2Data[i].leftIndex = -1
		}
		fmt.Scan(&tmp)
		if tmp != "-" {
			tree2Data[i].rightIndex, _ = strconv.Atoi(tmp)
		} else {
			tree2Data[i].rightIndex = -1
		}
	}
	tree2 := generateTree(tree2Data, tree2AIndex)
	fmt.Println(compare(tree1, tree2))
}

func compare(t1 *tree.Node, t2 *tree.Node) bool {

}

// 用slice的数据创建二叉树,index是返回的node(应传入A节点的下标)
func generateTree(s []nodeData, index int) *tree.Node {
	tmp := tree.CreateNode(s[index].data, nil, nil)
	if s[index].leftIndex != -1 {
		tmp.Left = generateTree(s, s[index].leftIndex)
	}
	if s[index].rightIndex != -1 {
		tmp.Right = generateTree(s, s[index].rightIndex)
	}
	return tmp
}
