package main

type MinStack struct {
	stack []int
	min   []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.stack = []int{x}
		this.min = []int{x}
	} else {
		if x < this.min[0] {
			this.min = append([]int{x}, this.min...)
		} else {
			this.min = append([]int{this.min[0]}, this.min...)
		}
		this.stack = append([]int{x}, this.stack...)
	}
}

func (this *MinStack) Pop() {
	this.stack = this.stack[1:]
	this.min = this.min[1:]
}

func (this *MinStack) Top() int {
	result := this.stack[0]
	return result
}

func (this *MinStack) Min() int {
	result := this.min[0]
	return result
}

func main() {
	stack := Constructor()
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)
}
