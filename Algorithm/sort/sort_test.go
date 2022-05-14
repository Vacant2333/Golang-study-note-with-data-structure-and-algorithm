package sort

import (
	"fmt"
	"testing"
)

// 测试用例
var a1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}
var a2 = []int{1, 2, 3, 4, 5, 6, 2, 7, 8, 9}
var a3 = []int{6, 8, 9, 7, 5, 4, 1, 2, 3, 9}
var a4 = []int{6, 8, 9, 7, 5, 4, 1, 2, 3, 9}
var a5 = []int{6, 8, 9, 7, 5, 4, 1, 2, 3, 9}

func TestIsSorted(t *testing.T) {
	fmt.Println(IsSorted(a1))
	fmt.Println(IsSorted(a2))
}

func TestIntBinarySearch(t *testing.T) {
	for index, v := range a1 {
		searchIndex, searchError := IntBinarySearch(a1, v)
		fmt.Println(index, v, searchIndex, searchError)
	}
	fmt.Println(IntBinarySearch(a1, 20))
}

func TestBubbleSort(t *testing.T) {
	BubbleSort(a3)
	fmt.Println(a3)
}

func TestInsertionSort(t *testing.T) {
	InsertionSort(a4)
	fmt.Println(a4)
}

func TestShellSort(t *testing.T) {
	ShellSort(a5)
	fmt.Println(a5)
}
