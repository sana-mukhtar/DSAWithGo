package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

func NewDoublyNode(data int) *Node {
	return &Node{
		data: data,
	}
}

// The head stores the pointer to our first node, and the tail stores the pointer to the last node. The length attribute stores the length of our linked list.
type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

// insertAtHead accepts data to add at the beginning of the list.
func (dll *DoublyLinkedList) insertAtHead(data int) {
	newNode := &Node{
		data: data,
		prev: nil,
		next: dll.head,
	}

	// If the list is empty, both head and tail point to the new node.
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		// 1. Tell the old head that a new node is now in front of it.
		dll.head.prev = newNode
		// 2. Move head pointer to this new node.
		dll.head = newNode
	}

	dll.length++
}

// insertAtHead accepts data to add at the end of the list.
func (dll *DoublyLinkedList) insertAtTail(data int) {
	newNode := &Node{
		data: data,
		prev: dll.tail,
		next: nil,
	}

	// If the list is empty, both head and tail point to the new node.
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		// 1. Tell the old tail that a new node is now in next of it.
		dll.tail.next = newNode
		// 2. Move tail pointer to this new node.
		dll.tail = newNode
	}

	dll.length++
}

// printLinkedlist shows the list as a chain of nodes.
func (dll *DoublyLinkedList) printLinkedlist() {
	current := dll.head
	if current == nil {
		fmt.Println("List is empty")
		return
	}

	fmt.Print("HEAD <-> ")
	for current != nil {
		fmt.Printf("[%d]", current.data)
		if current.next != nil {
			fmt.Print(" <-> ")
		}
		current = current.next
	}
	fmt.Println(" <-> TAIL")
}

func main() {
	dll := &DoublyLinkedList{}

	dll.insertAtHead(10)
	dll.insertAtHead(20)
	dll.insertAtHead(30)

	dll.printLinkedlist()
}
