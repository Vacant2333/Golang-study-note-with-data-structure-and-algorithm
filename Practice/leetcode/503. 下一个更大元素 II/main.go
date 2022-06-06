package main

import "fmt"

// https://leetcode.cn/problems/next-greater-element-ii/

func nextGreaterElements(nums []int) []int {
	n := len(nums)
	re := make([]int, n, n)
	for i := range re {
		re[i] = -1
	}
	stack := make([]int, 0, n)
	for i := 0; i < n*2-1; i++ {
		// 移除小于等于当前值的元素
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i%n] {
			re[stack[len(stack)-1]] = nums[i%n]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i%n)
	}
	return re
}

func main() {
	// answer: [2,-1,2]
	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	// answer: [2,3,4,-1,4]
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))
}
