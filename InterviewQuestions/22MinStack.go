package main

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
	}
}

func (this *MinStack) push(num int) {
	this.stack = append(this.stack, num)
}

func (this *MinStack) pop() int {
	stackLen := len(this.stack)
	lastElement := this.stack[stackLen-1]
	this.stack = this.stack[:stackLen-1]
	return lastElement
}

func (this *MinStack) top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) getMin() int {

	if len(this.stack) == 0 {
		return 0
	}

	min := this.stack[len(this.stack)-1]

	for _, val := range this.stack {
		if val < min {
			min = val
		}
	}

	return min

}
