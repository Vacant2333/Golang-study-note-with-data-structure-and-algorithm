package main

import (
	"fmt"
	"sort"
)

/*
https://pintia.cn/problem-sets/1497448825169559552/problems/1518544498597933056
给定N个（长整型范围内的）整数，要求输出从小到大排序后的结果。
本题旨在测试各种不同的排序算法在各种数据情况下的表现。各组测试数据特点如下：
数据1：只有1个元素；
数据2：11个不相同的整数，测试基本正确性；
数据3：103个随机整数；
数据4：104个随机整数；
数据5：105个随机整数；
数据6：105个顺序整数；
数据7：105个逆序整数；
数据8：105个基本有序的整数；
数据9：105个随机正整数，每个数字不超过1000。
输入格式:
输入第一行给出正整数N（≤10
5
 ），随后一行给出N个（长整型范围内的）整数，其间以空格分隔。

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
	//BubbleSort(s)
	sort.Ints(s)
	for i := 0; i < count; i++ {
		fmt.Print(s[i])
		if i != count-1 {
			fmt.Print(" ")
		}
	}
}

func BubbleSort(s []int) {
	for right := 1; right < len(s); right++ {
		// 如果一趟排序中flag没有变为true就表示数据已排序好,直接break
		flag := false
		// 右边界每次都会-1,因为每一趟排序之后,最大的元素就会被放到最后
		for left := 0; left < len(s)-right; left++ {
			if s[left] > s[left+1] {
				s[left], s[left+1] = s[left+1], s[left]
				flag = true
			}
		}
		if flag == false {
			break
		}
	}
}
