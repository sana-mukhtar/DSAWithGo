package main

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
		// Connect the current head’s prev pointer back to the new node.
		// If there are already boxes(nodes):
		// 1. Tell the old first box that a new box is now in front of it.
		dll.head.prev = newNode
		// Move the head pointer to the new node, making it the new first element.
		// 2. Move the "head" pointer to point at this new box, since it’s now the first box in the chain.
		dll.head = newNode
	}

	dll.length++
}
