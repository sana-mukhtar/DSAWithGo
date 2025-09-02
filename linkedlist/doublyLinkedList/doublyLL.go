package main

type DoublyNode struct {
	data int
	next *DoublyNode
	prev *DoublyNode
}

func NewDoublyNode(data int) *DoublyNode {
	return &DoublyNode{
		data: data,
	}
}
