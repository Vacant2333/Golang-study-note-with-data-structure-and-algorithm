package main

import "fmt"

/*
将一系列给定数字插入一个初始为空的小顶堆H[]。随后对任意给定的下标i，打印从H[i]到根结点的路径。
输入格式:
每组测试第1行包含2个正整数N和M(≤1000)，分别是插入元素的个数、以及需要打印的路径条数。下一行给出区间[-10000, 10000]内的N个要被插入一个初始为空的小顶堆的整数。最后一行给出M个下标。
输出格式:
对输入中给出的每个下标i，在一行中输出从H[i]到根结点的路径上的数据。数字间以1个空格分隔，行末不得有多余空格。
输入样例:
5 3
46 23 26 24 10
5 4 3
输出样例:
24 23 10
46 23 10
26 10
*/

type Heap []int

func main() {
	// 待插入元素个数, 需要打印的路径条数
	var elementCount, pathCount, tmp int
	heap := make(Heap, 0)
	fmt.Scan(&elementCount, &pathCount)
	// 读取元素
	for elementCount > 0 {
		fmt.Scan(&tmp)
		heap.Insert(tmp)
		elementCount--
		fmt.Println(heap)
	}
	pathList := make([]int, 0)
	// 读取要打印路径的节点
	for pathCount > 0 {
		fmt.Scan(&tmp)
		pathList = append(pathList, tmp)
		pathCount--
	}
}

// Insert 插入一个节点在Heap的尾部,然后Shift Up(上滤)最后一个节点,保证堆序性
func (heap *Heap) Insert(data int) {
	*heap = append(*heap, data)
	heap.shiftUp(len(*heap) - 1)
}

// shiftUp 上滤,将一个(最后一个)节点向上移动到正确的位置(最小堆)
func (heap *Heap) shiftUp(index int) {
	fmt.Println(index)
	for index > 0 {
		// 父节点
		fatherIndex := index / 2
		if (*heap)[fatherIndex] > (*heap)[index] {
			// 父节点比当前节点大,把当前节点和父节点交换
			heap.swap(fatherIndex, index)
			index /= 2
		} else {
			// 父节点比当前节点小 退出
			break
		}
	}
}

// swap 交换两个元素
func (heap *Heap) swap(index1, index2 int) {
	(*heap)[index1], (*heap)[index2] = (*heap)[index2], (*heap)[index1]
}
