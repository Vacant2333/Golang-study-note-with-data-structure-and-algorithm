package main

import "fmt"

/*
https://pintia.cn/problem-sets/1497448825169559552/problems/1500420206048747522
Given a stack which can keep M numbers at most. Push N numbers in the order of 1, 2, 3, ..., N and pop randomly.
You are supposed to tell if a given sequence of numbers is a possible pop sequence of the stack. For example, if M is 5 and N is 7,
we can obtain 1, 2, 3, 4, 5, 6, 7 from the stack, but not 3, 2, 1, 7, 5, 6, 4.
Input Specification:
Each input file contains one test case. For each case, the first line contains 3 numbers (all no more than 1000):
M (the maximum capacity of the stack), N (the length of push sequence), and K (the number of pop sequences to be checked).
Then K lines follow, each contains a pop sequence of N numbers. All the numbers in a line are separated by a space.
Output Specification:
For each pop sequence, print in one line "YES" if it is indeed a possible pop sequence of the stack, or "NO" if not.
Sample Input:
5 7 5
1 2 3 4 5 6 7
3 2 1 7 5 6 4
7 6 5 4 3 2 1
5 6 4 3 7 2 1
1 7 6 5 4 3 2
Sample Output:
YES
NO
NO
YES
NO
*/

func main() {
	var stackLength, sequenceLength, sequenceCount, tmp int
	// 栈最大长度,序列长度,待验证序列数量
	fmt.Scan(&stackLength, &sequenceLength, &sequenceCount)
	// 存待验证的序列
	sequenceMap := make(map[int][]int, 0)
	// 读入序列
	for i := 1; i <= sequenceCount; i++ {
		sequenceMap[i] = make([]int, 0)
		for j := 1; j <= sequenceLength; j++ {
			fmt.Scan(&tmp)
			sequenceMap[i] = append(sequenceMap[i], tmp)
		}
	}
	// 验证
	for i := 1; i <= sequenceCount; i++ {
		flag := true
		stack := initStack(stackLength, 0)
		inputStack := initStack(sequenceLength, sequenceLength)
		// 检查每一项
		for a := 0; a < sequenceLength; a++ {
			last, _ := stack.Pop(false)
			if last == sequenceMap[i][a] {
				// 这一项没问题
				stack.Pop(true)
			} else {
				// 最后一项不等于当前项,从inputStack中拿数据推入stack,直到相等,如果input被拿空了但仍需要就是错误的情况
				for {
					tmp, t := inputStack.Pop(true)
					if t == false {
						flag = false
						break
					} else {
						stack.Push(tmp)
					}
					tmp, _ = stack.Pop(false)
					if tmp == sequenceMap[i][a] {
						stack.Pop(true)
						break
					}
				}
			}
		}
		if flag {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

type stack struct {
	maxLength int
	data      []int
}

// initStack 初始化空stack
func initStack(maxLength, data int) *stack {
	stack := new(stack)
	stack.data = make([]int, 0)
	stack.maxLength = maxLength
	if data != 0 {
		for i := data; i >= 1; i-- {
			stack.data = append(stack.data, i)
		}
	}
	return stack
}

// Push 插入到尾部,满了返回false
func (s *stack) Push(i int) bool {
	if len(s.data) < s.maxLength {
		s.data = append(s.data, i)
		return true
	} else {
		return false
	}
}

// Pop 从尾部弹出一个,空的话返回false,delete代表是否删除它
func (s *stack) Pop(delete bool) (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	} else {
		lastIndex := len(s.data) - 1
		tmp := s.data[lastIndex]
		if delete {
			s.data = append(s.data[:lastIndex], s.data[lastIndex+1:]...)
		}
		return tmp, true
	}
}
