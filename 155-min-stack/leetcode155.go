package main

import "fmt"

type MinStack struct {
	stack []int
	min   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, 0)
		this.min = x
	} else {
		value := x - this.min
		if value < 0 {
			this.min = x
		}
		this.stack = append(this.stack, value)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	value := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if value < 0 {
		this.min -= value
	}
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		panic("invalid top operation on empty stack")
	}
	value := this.stack[len(this.stack)-1]
	if value > 0 {
		return this.min + value
	}
	return this.min
}

func (this *MinStack) GetMin() int {
	return this.min
}
