package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func createLinkedList(values []int) *Node {
	if len(values) == 0 {
		return nil
	}

	head := &Node{data: values[0]}
	current := head

	for _, val := range values[1:] {
		current.next = &Node{data: val}
		current = current.next
	}
	return head
}

func display(head *Node) {
	current := head

	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func reverse(head *Node) *Node {

	var prev *Node = nil
	current := head

	for current != nil {
		next := current.next
		current.next = prev
		prev = current
		current = next
	}

	return prev
}

func mergeSortedList(l1, l2 *Node) *Node {

	dummy := &Node{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.data <= l2.data {
			current.next = l1
			l1 = l1.next
		} else {
			current.next = l2
			l2 = l2.next
		}
		current = current.next
	}

	if l1 != nil {
		current.next = l1
	}

	if l2 != nil {
		current.next = l2
	}

	return dummy.next
}
