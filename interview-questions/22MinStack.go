package main

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

func (curr *MinStack) Push(val int) {
	curr.stack = append(curr.stack, val)
	if len(curr.minStack) == 0 || val <= curr.minStack[len(curr.minStack)-1] {
		curr.minStack = append(curr.minStack, val)
	}
}

func (curr *MinStack) Pop() {
	if len(curr.stack) == 0 {
		return
	}
	top := curr.stack[len(curr.stack)-1]
	curr.stack = curr.stack[:len(curr.stack)-1]

	if top == curr.minStack[len(curr.minStack)-1] {
		curr.minStack = curr.minStack[:len(curr.minStack)-1]
	}
}

func (curr *MinStack) Top() int {
	if len(curr.stack) == 0 {
		return -1
	}
	return curr.stack[len(curr.stack)-1]
}

func (curr *MinStack) GetMin() int {
	if len(curr.minStack) == 0 {
		return -1
	}
	return curr.minStack[len(curr.minStack)-1]
}
