package main

import "fmt"

// The head stores the pointer to our first node. The length attribute stores the length of our linked list.
type LinkedList struct {
	head   *Node
	length int
}

// Each node holds data and a pointer to the next node.
type Node struct {
	data int
	next *Node
}

func (ll *LinkedList) printLinkedlist() {
	curr := ll.head // start from head

	for curr != nil { // loop until the end of the list
		fmt.Print(curr.data) // print data
		if curr.next != nil {
			fmt.Print("->")
		}
		curr = curr.next
	}
}

// insertAtHead accepts data to add in linkedList chain. If the LL is empty it inserts the data to head otherwise store the head node to temporary variable and add the provided data to the head then points head's next node to the temporary head variable.
func (ll *LinkedList) insertAtHead(data int) {
	newNode := Node{
		data: data,
		next: nil,
	}

	if ll.head == nil {
		ll.head = &newNode
	} else {
		orgNode := ll.head
		ll.head = &newNode
		newNode.next = orgNode
	}
	ll.length++
}

// insertAtTail accepts data to add in linkedList chain. If the LL is empty it inserts the data to head, head and tail will be same, otherwise loop to the last node and point last nodes's next to the new node.
func (ll *LinkedList) insertAtTail(data int) {
	newNode := Node{
		data: data,
		next: nil,
	}

	if ll.head == nil {
		ll.head = &newNode
	} else {
		lastNode := ll.head // starts from head

		//
		for lastNode.next != nil {
			lastNode = lastNode.next
		}
		lastNode.next = &newNode
	}

	ll.length++
}

//  insert data at an arbitrary location.
func (ll *LinkedList) insert(position, data int) {
	if position == 0 {
		ll.insertAtHead(data)
		return
	}
	// 0 based index
	if position == ll.length-1 {
		ll.insertAtTail(data)
		return
	}

	newNode := &Node{
		data: data,
		next: nil,
	}

	justPrevNode := ll.head

	for i := 0; i < position-1; i++ {
		justPrevNode = justPrevNode.next
	}

	newNode.next = justPrevNode.next
	justPrevNode.next = newNode
	ll.length++
}

func (ll *LinkedList) deleteHead() {
	head := ll.head
	ll.head = head.next

	ll.length--
}

// deleteTail removes the last node of the linked list.
func (ll *LinkedList) deleteTail() {
	if ll.head == nil {
		return // empty list
	}

	if ll.head.next == nil {
		ll.head = nil // only one node
		ll.length--
		return
	}

	current := ll.head
	for current.next.next != nil {
		current = current.next
	}
	current.next = nil
	ll.length--
}

func (ll *LinkedList) delete(position int) {
	if position == 0 {
		ll.deleteHead()
	} else if position == ll.length-1 {
		ll.deleteTail()
	} else {
		prevNode := ll.head

		for i := 0; i < position-1; i++ {
			prevNode = prevNode.next
		}
		nodeToDelete := prevNode.next
		prevNode.next = nodeToDelete.next
		ll.length--
	}
}

// func reverseList(head *LinkedList) *LinkedList {
// 	var prev *LinkedList = nil
// 	curr := head

// 	for curr != nil {
// 		next := curr.Next // 1. Save next node
// 		curr.Next = prev  // 2. Reverse link
// 		prev = curr       // 3. Move prev forward
// 		curr = next       // 4. Move curr forward
// 	}

// 	return prev // prev is the new head

// }

func main() {
	// Create empty linked list
	list := LinkedList{}

	node1 := Node{data: 10}
	node2 := Node{data: 20}
	node3 := Node{data: 30}

	node1.next = &node2
	node2.next = &node3

	list.head = &node1
	list.length = 3

	list.printLinkedlist()

}
