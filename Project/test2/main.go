package main

import "fmt"

func main() {
	count := 0
	fmt.Scan(&count)
	for i := 0; i < count; i++ {
		// 总的节数
		c := 0
		fmt.Scan(&c)
		ok := true
		up := 0
		left := 0
		way := 1
		for a := 0; a < c; a++ {
			// 当前节 长度
			length := 0
			fmt.Scan(&length)
			switch way {
			case 1:
				//向右
				if length <= left {
					ok = false
				}
				left = length
			case 2:
				// 向下
				if length <= up {
					ok = false
				}
				up = length
			case 3:
				// 向左
				if length <= left {
					ok = false
				}
				left = length
			case 4:
				// 向上
				if length <= up {
					ok = false
				}
				up = length
				way = 0
			}
			way++
		}
		if ok {
			fmt.Println("NO")
		} else {
			fmt.Println("Yes")
		}
	}
}
