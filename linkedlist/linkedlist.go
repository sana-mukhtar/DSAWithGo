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

		fmt.Print(" ", curr.data, " ") // print data
		if curr.next != nil {
			fmt.Print("->")
		}
		curr = curr.next
	}

	fmt.Println()
}

// insertAtHead accepts data to add in linkedList chain.
//  If the LL is empty it inserts the data to head,
//  otherwise store the head node to temporary variable
//  and add the provided data to the head then
// points head's next node to the temporary head variable.
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

// insertAtTail accepts data to add in linkedList chain.
// If the LL is empty it inserts the data to head, head and tail will be same.
//  Otherwise loop to the last node and point last nodes's next to the new node.
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

	prevNode := ll.head

	for i := 0; i < position-1; i++ {
		prevNode = prevNode.next
	}

	newNode.next = prevNode.next
	prevNode.next = newNode
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
		prevNode.next = prevNode.next.next
		ll.length--
	}
}

// 10 -> 20 -> 30 -> nil
func (ll *LinkedList) reverseLinkedList() {
	var prev *Node // previous coach (starts empty)
	curr := ll.head

	for curr != nil {

		// reverse logic
		next := curr.next // save next
		curr.next = prev  // say current's node's next is previous itself

		// move forward
		prev = curr // now previous is old current
		curr = next // and current becomes the next
	}

	// head still points to 10, so we say head is last node i,e, 30 now
	ll.head = prev
}

// Brute Force, count total then return mid node
func (ll *LinkedList) findMiddleNode(head *Node) *Node {
	if head == nil {
		return head
	}

	curr := head
	count := 0
	for curr != nil {
		curr = curr.next
		count++
	}

	mid := count / 2
	midNode := head
	for i := 0; i < mid; i++ {
		midNode = midNode.next
	}

	return midNode
}

// Slow Fast Approach, slow moves 1 and fast moves 2 at a time
// slow will reach till mid when fast completes the traversal, return slow
func findMiddleNodeOptimal(head *Node) *Node {
	if head == nil {
		return head
	}

	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func mergeTwoLists(list1 *Node, list2 *Node) *Node {
	if list1 == nil && list2 == nil {
		return nil
	}
	head := &Node{}
	mergedList := head
	for list1 != nil && list2 != nil {
		if list1.data <= list2.data {
			mergedList = list1
			list1 = list1.next
		} else {
			mergedList = list2
			list2 = list2.next
		}
		mergedList = mergedList.next
	}

	if list1 != nil {
		mergedList.next = list1
	}
	if list2 != nil {
		mergedList.next = list2
	}

	return head.next
}

func main() {
	// Create empty linked list
	list := LinkedList{}

	node1 := Node{data: 10}
	node2 := Node{data: 20}
	node3 := Node{data: 30}
	node4 := Node{data: 40}
	node5 := Node{data: 50}

	node1.next = &node2
	node2.next = &node3
	node3.next = &node4
	node4.next = &node5

	list.head = &node1
	list.length = 3

	list.printLinkedlist()
	mid := list.findMiddleNode(&node2)
	fmt.Println(mid)
}
