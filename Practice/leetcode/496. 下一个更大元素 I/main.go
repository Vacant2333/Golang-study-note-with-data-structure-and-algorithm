package main

import "fmt"

// https://leetcode.cn/problems/next-greater-element-i/
// 单调栈+哈希表

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	// 单调栈(递增,先进去的元素大,后进去的小)
	stack := make([]int, 0)
	for i := len(nums2) - 1; i >= 0; i-- {
		// 去除不符合规律的栈中元素
		for len(stack) > 0 && nums2[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			m[nums2[i]] = stack[len(stack)-1]
		} else {
			m[nums2[i]] = -1
		}
		stack = append(stack, nums2[i])
	}
	for i := 0; i < len(nums1); i++ {
		nums1[i] = m[nums1[i]]
	}
	return nums1
}

func main() {
	// [-1,3,-1]
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
