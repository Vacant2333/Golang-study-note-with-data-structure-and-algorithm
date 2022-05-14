package main

import "fmt"

/*
https://pintia.cn/problem-sets/1497448825169559552/problems/1521131202348670976
给定公司N名员工的工龄，要求按工龄增序输出每个工龄段有多少员工。
输入格式:
输入首先给出正整数N（≤10^5），即员工总人数；随后给出N个整数，即每个员工的工龄，范围在[0, 50]。
输出格式:
按工龄的递增顺序输出每个工龄的员工个数，格式为：“工龄:人数”。每项占一行。如果人数为0则不输出该项。
输入样例:
8
10 2 0 5 7 2 5 2
输出样例:
0:1
2:3
5:2
7:1
10:1
*/

func main() {
	var count int
	fmt.Scan(&count)
	data := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&data[i])
	}
	ShellSort(data)
	value := -1
	number := 0
	for i := 0; i < count; i++ {
		if value == data[i] {
			number++
		} else {
			// 不相等 输出
			if value != -1 {
				fmt.Printf("%d:%d\n", value, number)
			}
			value = data[i]
			number = 1
		}
	}
	fmt.Printf("%d:%d\n", value, number)
}

func ShellSort(s []int) {
	/*
		基本的希尔排序的增量序列是n,n/2,n/4...
		但是它的时间效率也是O(n^2)
		这里使用Sedgewick增量序列,效率更高
		他的时间效率是O(n^6/5),但是目前无法证明
		这里只写了Sedgewick的一部分增量
		https://baike.baidu.com/item/%E5%B8%8C%E5%B0%94%E5%A2%9E%E9%87%8F/6853339?fr=aladdin
	*/
	sedgewick := []int{929, 505, 209, 109, 41, 19, 5, 1, 0}
	sedgewickIndex := 0
	// 增量不能大于数组长度
	for sedgewick[sedgewickIndex] > len(s) {
		sedgewickIndex++
	}
	// 可以把插入排序的增量理解为1,这里换成sedgewick的元素,一直想下取直到0
	for d := sedgewick[sedgewickIndex]; d > 0; d = sedgewick[sedgewickIndex] {
		// 插入排序
		for i := d; i < len(s); i++ {
			tmp := s[i]
			k := i
			for ; k >= d && s[k-d] > tmp; k -= d {
				s[k] = s[k-d]
			}
			s[k] = tmp
		}
		// 增量取下一个
		sedgewickIndex++
	}
}
