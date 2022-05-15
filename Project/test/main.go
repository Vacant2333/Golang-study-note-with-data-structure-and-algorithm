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
	s := make([][]int, 2)
	s[0] = []int{15, 1}
	s[1] = []int{10, 11}
	// 2
	fmt.Println(maximumWhiteTiles(s, 2))
}

func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	QuickSort(0, len(tiles)-1, tiles)
	max := 0
	for i := 0; i < len(tiles); i++ {
		lenTmp := carpetLen - getLen(tiles[i])
		lenSum := getLen(tiles[i])
		nowIndex := i + 1
		// 拿到从这个下标开始的覆盖的长度
		for lenTmp > 0 {
			if nowIndex >= len(tiles) {
				break
			}
			// 处理空白
			lenTmp -= tiles[nowIndex][0] - tiles[nowIndex-1][1] - 1
			if lenTmp <= 0 || lenSum+lenTmp <= max {
				break
			}
			thisLen := getLen(tiles[nowIndex])
			if lenTmp >= thisLen {
				lenTmp -= thisLen
				lenSum += thisLen
				nowIndex++
			} else {
				// 剩下的长度比现在这条短
				lenSum += lenTmp
				break
			}
		}
		if lenSum > max {
			max = lenSum
		}
	}

	return max
}

func getLen(x []int) int {
	return x[1] - x[0] + 1
}

func QuickSort(left int, right int, array [][]int) {
	l := left
	r := right
	pivot := array[(left+right)/2]
	var temp []int
	for l < r {
		for array[l][0] < pivot[0] {
			l++
		}
		for array[r][0] > pivot[0] {
			r--
		}
		if l >= r {
			break
		}
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		if array[l][0] == pivot[0] {
			r--
		}
		if array[r][0] == pivot[0] {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(left, r, array)
	}
	if right > l {
		QuickSort(l, right, array)
	}
}
