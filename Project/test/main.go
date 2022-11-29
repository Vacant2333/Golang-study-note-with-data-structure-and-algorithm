package main

import (
	"fmt"
	"hash/fnv"
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
	// fmt.Println(unicode.IsLetter('A'))
	// fmt.Println(os.Getuid())
	//fmt.Println(ihash("alloy"), ihash("work"), ihash("out"))
	//
	//defer func() {
	//	fmt.Println(1)
	//}()
	//defer func() {
	//	a := 0
	//	fmt.Println(2 / a)
	//}()

	//sendData := make(map[string]interface{})
	//
	//sendData["topcolor"] = "#FF0000"
	//data := make(map[string]map[string]interface{})
	//data["nowDate"] = make(map[string]interface{})
	//data["nowDate"]["value"] = "黄先生"
	//
	//sendData["data"] = data

	a := []int{1, 2, 3}
	change(&a)
	fmt.Println(a)
}

func change(a *[]int) {
	s := *a
	s[0] = 0

	a = &s
}

func ihash(key string) int {
	h := fnv.New32a()
	h.Write([]byte(key))
	return int(h.Sum32() & 0x7fffffff)
}
