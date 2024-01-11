package stack

import "math"

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minStack) == 0 {
		this.minStack = append(this.minStack, val)
	} else {
		top := this.minStack[len(this.minStack)-1]
		this.minStack = append(this.minStack, minInt(top, val))
	}

}

func minInt(i, j int) int {
	if i <= j {
		return i
	}
	return j
}

func (this *MinStack) Pop() {
	if len(this.stack) >= 1 {
		this.stack = this.stack[:len(this.stack)-1]
		this.minStack = this.minStack[:len(this.minStack)-1]
	}

}

func (this *MinStack) Top() int {
	if len(this.stack) >= 1 {
		return this.stack[len(this.stack)-1]
	}
	return math.MaxInt
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) >= 1 {
		return this.minStack[len(this.minStack)-1]
	}
	return math.MaxInt
}
