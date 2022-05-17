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

type CountIntervals struct {
	data [][2]int
}

func Constructor() CountIntervals {
	data := make([][2]int, 0, 100000)
	return CountIntervals{data}
}

func (this *CountIntervals) Add(left int, right int) {
	this.data = append(this.data, [2]int{left, right})
}

func (this *CountIntervals) Count() int {
	this.f()
	count := 0
	for _, v := range this.data {
		count += v[1] - v[0] + 1
	}
	return count
}

// 优化
func (this *CountIntervals) f() {
	fmt.Println(this.data)
	if len(this.data) > 0 {
		QuickSort(0, len(this.data)-1, this.data)
		k := 0
		tmp := [2]int{0, 0}
		n := make([][2]int, 0, 100000)
		for k < len(this.data) {
			fmt.Println("k", k)
			fmt.Println("kd", this.data[k])
			if tmp[0] == 0 || tmp[1] >= this.data[k][0] {
				if tmp[0] == 0 {
					tmp[0] = this.data[k][0]
					tmp[1] = this.data[k][1]
				} else {
					if tmp[1] < this.data[k][1] {
						// 不包含下一个区间
						tmp[1] = this.data[k][1]
					}
				}
				fmt.Println("kk", this.data[k])
			} else {
				// 不包含下一个
				n = append(n, tmp)
				fmt.Println("kkk", tmp)
				tmp[0] = this.data[k][0]
				tmp[1] = this.data[k][1]

			}
			k++
		}
		if tmp[0] != 0 {
			fmt.Println("tttttt", tmp)
			n = append(n, tmp)
		}
		this.data = n[:]
		fmt.Println(n)
	}
}

func QuickSort(left int, right int, array [][2]int) {
	l := left
	r := right
	pivot := array[(left+right)/2]
	var temp [2]int
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
