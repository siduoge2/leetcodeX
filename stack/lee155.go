package stack

import (
	"math"
)

type MinStack struct {
	stck []int
}

func Constructor() MinStack {
	return MinStack{stck: nil}
}

func (this *MinStack) Push(i int) {
	this.stck = append(this.stck, i)
}

func (this *MinStack) isEmpty() bool {
	l := len(this.stck)
	if l < 1 {
		return true
	}
	return false
}

func (this *MinStack) Pop() {
	if this.isEmpty() {
		return
	}
	l := len(this.stck)
	this.stck = this.stck[:l-1]
}

func (this *MinStack) Top() int {
	if this.isEmpty() {
		return math.MaxInt8
	}
	l := len(this.stck)
	return this.stck[l-1]
}

func (this *MinStack) GetMin() int {
	min := math.MaxInt64
	if this.isEmpty() {
		return min
	}
	for _, v := range this.stck {
		if v < min {
			min = v
		}
	}
	return min
}
