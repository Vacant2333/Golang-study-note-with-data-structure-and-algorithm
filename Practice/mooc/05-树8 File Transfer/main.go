package main

import (
	"fmt"
	"strconv"
)

/*
We have a network of computers and a list of bidirectional connections. Each of these connections allows a file transfer from one computer to another. Is it possible to send a file from any computer on the network to any other?
Input Specification:
Each input file contains one test case. For each test case, the first line contains N (2≤N≤104), the total number of computers in a network. Each computer in the network is then represented by a positive integer between 1 and N. Then in the following lines, the input is given in the format:
I c1 c2
where I stands for inputting a connection between c1 and c2; or
C c1 c2
where C stands for checking if it is possible to transfer files between c1 and c2; or
S
where S stands for stopping this case.
Output Specification:
For each C case, print in one line the word "yes" or "no" if it is possible or impossible to transfer files between c1 and c2,
respectively. At the end of each case, print in one line "The network is connected." if there is a path between any pair of computers;
or "There are k components." where k is the number of connected components in this network.
Sample Input 1:
5
C 3 2
I 3 2
C 1 5
I 4 5
I 2 4
C 3 5
S
Sample Output 1:
no
no
yes
There are 2 components.
Sample Input 2:
5
C 3 2
I 3 2
C 1 5
I 4 5
I 2 4
C 3 5
I 1 3
C 1 5
S
Sample Output 2:
no
no
yes
yes
The network is connected.
*/

func main() {
	// 元素个数
	var n int
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		SetData = append(SetData, Node{ElementType(i), -1})
	}
	var command string
	var c1, c2 int
	for {
		fmt.Scan(&command)
		if command != "S" {
			fmt.Scan(&c1, &c2)
			if command == "C" {
				// 检查两台电脑的连接
				if Find(ElementType(c1)) == Find(ElementType(c2)) {
					fmt.Println("yes")
				} else {
					fmt.Println("no")
				}
			} else if command == "I" {
				// 连接两台电脑
				Union(Find(ElementType(c1)), Find(ElementType(c2)))
			}
		} else {
			// S指令 退出
			break
		}
	}
	// 检查有几个节点的父节点是-1 如果是0个则全都连起来了
	fatherNodeCount := 0
	for _, v := range SetData {
		if v.Father < 0 {
			fatherNodeCount++
		} else {
			//fmt.Println(v)
		}
	}
	if fatherNodeCount <= 1 {
		fmt.Println("The network is connected.")
	} else {
		fmt.Println("There are " + strconv.Itoa(fatherNodeCount) + " components.")
	}
}

// ElementType 集合的数据类型
type ElementType int

// SetName 集合的名称,根节点的下标作为名称
type SetName int

type Set []Node

type Node struct {
	Data   ElementType
	Father SetName
}

// SetData 存储所有的集合的节点
var SetData Set

// Union 合并两个集合(将较小的集合的所有元素的Father改为较大的集合的Father),返回新集合的名称
func Union(set1, set2 SetName) SetName {
	if SetData[set1].Father*-1 >= SetData[set2].Father*-1 {
		// 更新节点总数
		SetData[set1].Father += SetData[set2].Father
		// set1的元素比set2多，把set2的元素的Father改为set1
		SetData[set2].Father = set1
		for i, v := range SetData {
			if v.Father == set2 {
				SetData[i].Father = set1
			}
		}
		return set1
	} else {
		SetData[set2].Father += SetData[set1].Father
		SetData[set1].Father = set2
		for i, v := range SetData {
			if v.Father == set1 {
				SetData[i].Father = set2
			}
		}
		return set2
	}
}

// Find 在集合中查找一个节点，如果没有返回-1，有的话返回父节点的下标
func Find(data ElementType) SetName {
	for i := 0; i < len(SetData); i++ {
		if SetData[i].Data == data {
			if SetData[i].Father >= 0 {
				// 它是一个子节点 返回父节点
				return SetData[i].Father
			} else {
				// 它是一个父节点 直接返回
				return SetName(i)
			}
		}
	}
	return -1
}
