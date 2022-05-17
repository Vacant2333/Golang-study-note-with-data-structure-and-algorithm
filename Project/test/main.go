package main

import "fmt"

// fibonacci 斐波那契 返回一个“返回int的函数”
func fibonacci() func() int {
	s := []int{0, 1}
	n := 0
	return func() int {
		s = append(s, s[len(s)-2]+s[len(s)-1])
		n++
		return s[n-1]
	}
}

func main() {
	/*
		f := fibonacci()
		for i := 0; i < 10; i++ {
			fmt.Println(f())
		}
		t := tree.CreateNode("1", nil, nil)
		fmt.Println(t)
		s := make([]int, 5)
		fmt.Println(s[0], &s[0])
	*/
	c := Constructor()
	c.Add(2, 3)
	c.Add(7, 10)
	fmt.Println(c.Count())
	fmt.Println(c.data)
	// 6

}
