package main

import "fmt"

func main() {
	count := 0
	fmt.Scanf("%d", &count)
	for i := 0; i < count; i++ {
		// 车数量
		carCount := 0
		fmt.Scanf("%d", &carCount)
		stack := make([]int, 0)
		push := make([]int, carCount)
		pop := make([]int, carCount)
		for i := 0; i < carCount; i++ {
			fmt.Scanf("%d", &push[i])
		}
		for i := 0; i < carCount; i++ {
			fmt.Scanf("%d", &pop[i])
		}
		c := 0
		for _, x := range push {
			stack = append(stack, x)
			for len(stack) > 0 && stack[len(stack)-1] == pop[c] {
				stack = stack[:len(stack)-1]
				c++
			}
		}
		if len(stack) == 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
