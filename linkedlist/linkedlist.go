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

func (ll *LinkedList) insertAtTail(data int) {
	newNode := Node{
		data: data,
		next: nil,
	}

	if ll.head == nil {
		ll.head = &newNode
	} else {
		existingNode := ll.head
		for existingNode.next != nil {
			existingNode = existingNode.next
		}
		existingNode.next = &newNode
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

func main() {
	// We can initialize our linked list this way. The initial length is zero because there are no nodes in our list yet. The head points to nil because there are no nodes to point to.
	list := LinkedList{nil, 0}

	for i := 0; i < 5; i++ {
		if i == 0 {
			fmt.Println(list)
		}
		list.insertAtHead(i)
		fmt.Println(list)
	}

}
