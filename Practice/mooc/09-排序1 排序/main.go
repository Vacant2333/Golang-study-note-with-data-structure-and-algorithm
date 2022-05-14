package main

import (
	"fmt"
)

/*
https://pintia.cn/problem-sets/1497448825169559552/problems/1518544498597933056
给定N个（长整型范围内的）整数，要求输出从小到大排序后的结果。
本题旨在测试各种不同的排序算法在各种数据情况下的表现。各组测试数据特点如下：
数据1：只有1个元素；
数据2：11个不相同的整数，测试基本正确性；
数据3：10^3个随机整数；
数据4：10^4个随机整数；
数据5：10^5个随机整数；
数据6：10^5个顺序整数；
数据7：10^5个逆序整数；
数据8：10^5个基本有序的整数；
数据9：10^5个随机正整数，每个数字不超过1000。
输入格式:
输入第一行给出正整数N（≤10^5），随后一行给出N个（长整型范围内的）整数，其间以空格分隔。
输出格式:
在一行中输出从小到大排序后的结果，数字间以1个空格分隔，行末不得有多余空格。
输入样例:
11
4 981 10 -17 0 -20 29 50 8 43 -5
输出样例:
-20 -17 -5 0 4 8 10 29 43 50 981
*/

func main() {
	var count int
	fmt.Scan(&count)
	s := make([]int, count)
	for i := 0; i < count; i++ {
		fmt.Scan(&s[i])
	}
	// 超时了=.=
	ShellSort(s)
	//sort.Ints(s)
	for i := 0; i < count; i++ {
		fmt.Print(s[i])
		if i != count-1 {
			fmt.Print(" ")
		}
	}
}

// ShellSort 希尔排序
func ShellSort(s []int) {
	/*
		基本的希尔排序的增量序列是n,n/2,n/4...
		但是它的时间效率也是O(n^2)
		这里使用Sedgewick增量序列,效率更高
		他的时间效率是O(n^6/5),但是目前无法证明
		这里只写了Sedgewick的一部分增量
	*/
	sedgewick := []int{929, 505, 209, 109, 41, 19, 5, 1, 0}
	sedgewickIndex := 0
	// 增量不能大于数组长度
	for sedgewick[sedgewickIndex] > len(s) {
		sedgewickIndex++
	}
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
