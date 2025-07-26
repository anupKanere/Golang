package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
}

func (dl *DoublyLinkedList) InsertAtBegining(val int) {
	newNode := &Node{data: val}
	if dl.head == nil {
		dl.head = newNode
		return
	}

	newNode.next = dl.head
	dl.head.prev = newNode
	dl.head = newNode

}

func (dl *DoublyLinkedList) InsertAtEnd(val int) {
	newNode := &Node{data: val}
	if dl.head == nil {
		dl.head = newNode
		return
	}

	curr := dl.head
	for curr.next != nil {
		curr = curr.next
	}
	newNode.prev = curr
	curr.next = newNode
}

func (dl *DoublyLinkedList) Delete(val int) {
	if dl.head == nil {
		return
	}

	if dl.head.data == val {
		dl.head = dl.head.next
		if dl.head != nil {
			dl.head.prev = nil
		}
		return
	}

	curr := dl.head
	for curr != nil && curr.data != val {
		curr = curr.next
	}

	if curr == nil {
		fmt.Println("Node not Present")
		return
	}

	if curr.next != nil {
		curr.next.prev = curr.prev
	}

	if curr.prev != nil {
		curr.prev.next = curr.next
	}

}

func (dl *DoublyLinkedList) DisplayForward() {
	if dl.head == nil {
		fmt.Println("Doubly  Linked List is Empty!!!")
		return
	}
	fmt.Println()
	fmt.Print("PRINTING FOWARD  DOUBLY LINKED LIST : ")
	curr := dl.head
	for curr != nil {
		fmt.Printf("%d -> ", curr.data)
		curr = curr.next
	}
	fmt.Printf("Nil")
}

func (dl *DoublyLinkedList) DisplayReverse() {
	if dl.head == nil {
		fmt.Println("DOUBLY LINKED LIST IS EMPTY!!!")
		return
	}

	curr := dl.head
	for curr.next != nil {
		curr = curr.next
	}

	fmt.Println()
	fmt.Print("PRINTING  REVERSE DOUBLY LINKED LIST : ")
	for curr != nil {
		fmt.Printf("%d -> ", curr.data)
		curr = curr.prev
	}

	fmt.Printf("Nil")
}
func main() {

	dll := &DoublyLinkedList{}
	dll.DisplayForward()
	dll.DisplayReverse()
	dll.InsertAtBegining(40)
	dll.InsertAtBegining(30)
	dll.InsertAtBegining(20)
	dll.InsertAtBegining(10)
	dll.InsertAtEnd(50)
	dll.DisplayForward()
	dll.DisplayReverse()
	dll.Delete(20)
	dll.DisplayForward()
	dll.DisplayReverse()

}
