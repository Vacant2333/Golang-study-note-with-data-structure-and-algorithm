package heap

import (
	"testing"
)

var heap Heap
var n = 1000000

func TestCreate(t *testing.T) {
	heap = Create()
}

func TestHeap_Insert(t *testing.T) {
	for i := 0; i <= n; i++ {
		heap.Insert(ElementType(i))
	}
}

func TestHeap_GetMax(t *testing.T) {
	if heap.validate() == false {
		t.Errorf("validate error!")
	}
	for i := n; i >= 0; i-- {
		if heap.GetMax() != ElementType(i) {
			t.Errorf("GetMax error!")
		}
	}
}
