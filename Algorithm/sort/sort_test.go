package sort

import (
	"fmt"
	"testing"
)

var a1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func TestIsSorted(t *testing.T) {
	fmt.Println(IsSorted(IntSlice(a1)))
}
