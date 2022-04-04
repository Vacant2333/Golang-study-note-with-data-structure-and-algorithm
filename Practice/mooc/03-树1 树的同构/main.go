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
	// 检查是否为空
	if tree1 == nil || tree2 == nil {
		if tree1 == nil && tree2 == nil {
			fmt.Print("Yes")
		} else {
			fmt.Print("No")
		}
		return
	}
	if compare(tree1, tree2) {
		fmt.Print("Yes")
	} else {
		fmt.Print("No")
	}
}

func compare(t1 *tree.Node, t2 *tree.Node) bool {
	if t1.CountSon() == t2.CountSon() {
		// 子节点数量相同
		if t1.CountSon() == 0 {
			// 都没有节点
			return t1.Data == t2.Data
		}
		if t1.CountSon() == 1 {
			// 都只有一个节点 找出他们的两个要比较的节点
			var tmp1, tmp2 *tree.Node
			if t1.Left != nil {
				tmp1 = t1.Left
			} else {
				tmp1 = t1.Right
			}
			if t2.Left != nil {
				tmp2 = t2.Left
			} else {
				tmp2 = t2.Right
			}
			// 比较两个节点
			return tmp1.Data == tmp2.Data && compare(tmp1, tmp2)
		}
		if t1.CountSon() == 2 {
			// 都有两个节点
			if t1.Left.Data == t2.Left.Data {
				// 没有交换,直接比较同向节点
				return t1.Left.Data == t2.Left.Data && compare(t1.Left, t2.Left) && t1.Right.Data == t2.Right.Data && compare(t1.Right, t2.Right)
			} else if t1.Left.Data == t2.Right.Data {
				// 有交换
				return t1.Left.Data == t2.Right.Data && compare(t1.Left, t2.Right) && t1.Right.Data == t2.Left.Data && compare(t1.Right, t2.Left)
			} else {
				// 完全比不了
				return false
			}
		}
	} else {
		// 节点数量不同
		return false
	}
	return false
}

// 用slice的数据创建二叉树,index是返回的node(应传入A节点的下标)
func generateTree(s []nodeData, index int) *tree.Node {
	// 是否为空
	if len(s) == 0 {
		return nil
	}
	tmp := tree.CreateNode(s[index].data, nil, nil)
	if s[index].leftIndex != -1 {
		tmp.Left = generateTree(s, s[index].leftIndex)
	}
	if s[index].rightIndex != -1 {
		tmp.Right = generateTree(s, s[index].rightIndex)
	}
	return tmp
}
