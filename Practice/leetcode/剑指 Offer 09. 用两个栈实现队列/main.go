package main

import "fmt"

type CQueue struct {
	left  []int
	right []int
}

func Constructor() CQueue {
	return CQueue{
		left:  make([]int, 0),
		right: make([]int, 0),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.put("left", value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.right) == 0 {
		if len(this.left) == 0 {
			return -1
		}
		// 将left的内容移入right，然后拿到right的第一个值
		for {
			value := this.get("left")
			if value == -1 {
				break
			}
			this.put("right", value)
		}
	}
	return this.get("right")
}

func (this *CQueue) put(stack string, value int) {
	if stack == "left" {
		this.left = append([]int{value}, this.left...)
	} else {
		this.right = append([]int{value}, this.right...)
	}
}

func (this *CQueue) get(stack string) int {
	if stack == "left" {
		if len(this.left) == 0 {
			return -1
		}
		result := this.left[0]
		this.left = this.left[1:]
		return result
	} else {
		if len(this.right) == 0 {
			return -1
		}
		result := this.right[0]
		this.right = this.right[1:]
		return result
	}
}

func main() {
	queue := Constructor()
	queue.AppendTail(1)
	queue.AppendTail(2)
	fmt.Println(queue.DeleteHead(), queue.DeleteHead())
}
