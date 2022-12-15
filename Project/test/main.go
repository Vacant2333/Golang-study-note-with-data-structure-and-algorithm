package main

import "fmt"

func main() {
	a1 := []int{12, 1, 2, 3, 4, 5, 6, 14, 13}

	for _, v := range a1 {
		a1 = []int{1}
		fmt.Println(v)
	}
}
