package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l1 *LinkedList) Add(val int) {
	fmt.Println("Val - ", val)
	newNode := &Node{data: val}
	if l1.head == nil {
		fmt.Println("first node to be added")
		l1.head = newNode
		return
	}

	curr := l1.head

	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode

}

func (l1 *LinkedList) PrintLinkedList() {

	fmt.Println()
	if l1.head == nil {
		fmt.Println("Linked List is Empty!!!")
		return
	}

	// fmt.Println()
	curr := l1.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.data)
		curr = curr.next
	}
	fmt.Printf("nil\n")
}

func (l1 *LinkedList) Delete(val int) {
	if l1.head == nil {
		return
	}

	if l1.head.data == val {
		l1.head = l1.head.next
		return
	}

	curr := l1.head.next
	prev := l1.head

	for curr != nil {
		if curr.data == val {
			prev.next = curr.next
		}
		prev = curr
		curr = curr.next
	}
}

func (l1 *LinkedList) Search(val int) bool {
	if l1.head == nil {
		return false
	}

	curr := l1.head
	for curr != nil {
		if curr.data == val {
			return true
		}
		curr = curr.next
	}
	return false
}

func (l1 *LinkedList) ReverseLinkedList() {
	var prev *Node
	curr := l1.head
	var next *Node

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l1.head = prev
}

func main() {

	l1 := LinkedList{}
	l1.Add(10)
	l1.Add(20)
	l1.Add(30)
	l1.PrintLinkedList()
	l1.Delete(10)
	// l1.Delete(20)
	// l1.Delete(30)
	l1.ReverseLinkedList()
	l1.PrintLinkedList()
	fmt.Println("Found : ", l1.Search(20))
}
