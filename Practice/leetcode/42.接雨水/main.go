package main

/*
https://leetcode.cn/problems/trapping-rain-water/
DP
*/

func trap(height []int) int {
	L, sum := len(height), 0
	if L == 0 {
		return 0
	}
	leftMax := make([]int, L, L)
	rightMax := make([]int, L, L)

	left := 0
	for index := 0; index < L; index++ {
		leftMax[index] = left
		left = max(left, height[index])
	}
	right := 0
	for index := L - 1; index >= 0; index-- {
		rightMax[index] = right
		right = max(right, height[index])
	}

	for index := 0; index < L; index++ {
		sum += max(min(leftMax[index], rightMax[index])-height[index], 0)
	}
	return sum
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
