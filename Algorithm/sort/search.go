package sort

import (
	"errors"
)

// IntBinarySearch BinarySearch 二分查找(s必须是排序好的(升序)) 返回下标
func IntBinarySearch(s []int, value int) (int, error) {
	// 没有排序的话返回错误
	if IsSorted(IntSlice(s)) == false {
		return -1, errors.New("data is not sorted")
	}
	left, right := 0, len(s)
	for left < right {
		mid := (left + right) / 2
		if s[mid] > value {
			// value在mid的左边,将right向左移动,到mid-1
			right = mid - 1
		} else if s[mid] < value {
			left = mid + 1
		} else {
			return mid, nil
		}
	}
	return -1, nil
}
