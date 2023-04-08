package Template

import (
	"fmt"
	"testing"
)

func TestGetSumN(t *testing.T) {
	testCases := map[int]int{
		0: 0,
		1: 1,
		6: 21,
		8: 36,
	}
	for n, sum := range testCases {
		fmt.Println(n, GetSumN(n), sum == GetSumN(n))
	}
}
