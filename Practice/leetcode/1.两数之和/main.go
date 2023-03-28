package main

/*
https://leetcode.cn/problems/two-sum/
*/

//func twoSum(nums []int, target int) []int {
//	sort.Ints(nums)
//	L := len(nums)
//	left := 0
//	right := L - 1
//
//	for left != right {
//		sum := nums[left] + nums[right]
//		if sum == target {
//			return []int{nums[left], nums[right]}
//		} else if sum > target {
//			right--
//		} else {
//			left++
//		}
//	}
//	return []int{}
//}

func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for index1, value1 := range nums {
		if index2, ok := m[target-value1]; ok {
			return []int{index1, index2}
		}
		m[value1] = index1
	}
	return []int{}
}
