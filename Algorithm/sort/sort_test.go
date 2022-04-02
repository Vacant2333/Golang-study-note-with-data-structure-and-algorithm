package sort

import (
	"fmt"
	"testing"
)

// 测试用例
var a1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}
var a2 = []int{1, 2, 3, 4, 5, 6, 2, 7, 8, 9}

func TestIsSorted(t *testing.T) {
	fmt.Println(IsSorted(IntSlice(a1)))
	fmt.Println(IsSorted(IntSlice(a2)))
}

func TestBinarySearch(t *testing.T) {
	for index, v := range a1 {
		searchIndex, searchError := BinarySearch(a1, v)
		fmt.Println(index, v, searchIndex, searchError)
	}
	fmt.Println(BinarySearch(a1, 20))
}
