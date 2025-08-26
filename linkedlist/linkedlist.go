package main

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
