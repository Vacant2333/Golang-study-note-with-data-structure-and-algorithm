package main

import (
	"fmt"
	"math/rand"
)

/*
三门问题
我们有3个门,
*/

var testCount = 100000

func main() {
	// 换/不换 赢的次数
	t, f := 0, 0

	for i := 0; i <= testCount; i++ { // 随机选择一个门 填入true作为大奖
		price := rand3()
		// 选手随机选择一个门
		choose := rand3()
		if choose == price {
			// 选择的是大奖的门,不换就能赢
			f++
		} else {
			t++
		}
	}

	fmt.Printf("换:%d 不换:%d", t, f)

}

// 获得一个[0, 2]内的随机数
func rand3() int {
	return rand.Int() % 3
}
