package main

import (
	"fmt"
	"strconv"
)

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
	decode("1234567")
}

func decode(num string) {
	for i, v := range num {
		fmt.Printf("range: index:%v value:%v int32(v):%v \n", i, v, v&15)
	}
	fmt.Println("-----------")
	for i := 0; i < len(num); i++ {
		tmp, _ := strconv.Atoi(num[i])
		fmt.Printf("i index:%v int32(v):%v", i, num[i])
	}

}
