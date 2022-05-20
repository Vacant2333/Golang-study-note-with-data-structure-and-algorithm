package heap

import (
	"fmt"
	"testing"
)

var heap Heap

func TestCreate(t *testing.T) {
	heap = Create()
	fmt.Println(heap)
}

func TestHeap_Insert(t *testing.T) {
	heap = Create()
	heap.Insert(1)
	heap.Insert(3)
	heap.Insert(5)
	heap.Insert(6)
	heap.Insert(4)
	heap.Insert(7)
	fmt.Println(heap)
}

func TestHeap_GetMax(t *testing.T) {
	heap = Create()
	heap.Insert(1)
	heap.Insert(3)
	heap.Insert(5)
	heap.Insert(6)
	heap.Insert(4)
	heap.Insert(7)
	fmt.Println(heap)
	fmt.Println("Max: ", heap.GetMax())
	fmt.Println(heap)
	fmt.Println(heap.validate())
}
