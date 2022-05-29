package main

import "fmt"

// 没做出来

func totalSteps(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		tmp := 0
		lastIndex := i + 1
		for x := i + 1; x < len(nums); x++ {
			if nums[x] < nums[lastIndex] {
				break
			} else {
				lastIndex = x
			}

			if nums[x] < nums[i] {
				tmp++
			} else {
				break
			}
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func main() {
	fmt.Println(totalSteps([]int{10, 1, 2, 3, 4, 5, 6, 1, 2, 3}))
}
