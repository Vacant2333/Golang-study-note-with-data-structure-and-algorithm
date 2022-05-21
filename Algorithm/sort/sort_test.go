package sort

import (
	"math/rand"
	"testing"
)

// 测试用例
var s1, s2, s3, s4 []int

func init() {
	// 初始化测试数据,100w个元素
	for i := 0; i < 1000000; i++ {
		n := rand.Int()
		s1 = append(s1, n)
		s2 = append(s2, n)
		s3 = append(s3, n)
		s4 = append(s4, n)
	}
}

func TestBubbleSort(t *testing.T) {
	BubbleSort(s1)
	if IsSorted(s1) == false {
		t.Error("BubbleSort error")
	}
}

func TestInsertionSort(t *testing.T) {
	InsertionSort(s2)
	if IsSorted(s2) == false {
		t.Error("InsertionSort error")
	}
}

func TestShellSort(t *testing.T) {
	ShellSort(s3)
	if IsSorted(s3) == false {
		t.Error("ShellSort error")
	}
}

func TestHeapSort(t *testing.T) {
	HeapSort(s4)
	if IsSorted(s4) == false {
		t.Error("HeapSort error")
	}
}
