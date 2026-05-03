package main

import "fmt"

type stack struct {
	elements []int
}

func newStack() *stack {
	return &stack{
		elements: make([]int, 0),
	}
}

func (st *stack) push(elem int) {
	st.elements = append(st.elements, elem)
}

func (st *stack) pop() int {
	lastIdx := len(st.elements) - 1
	lastElem := st.elements[lastIdx]

	st.elements = st.elements[:lastIdx]

	return lastElem
}

func (st *stack) peek() int {
	lastIdx := len(st.elements) - 1
	lastElem := st.elements[lastIdx]

	return lastElem
}

func (st *stack) size() int {
	return len(st.elements)
}

func (st *stack) isEmpty() bool {
	return len(st.elements) == 0
}

func main() {
	st := newStack()
	st.push(1)
	st.push(2)
	st.push(3)
	st.push(4)
	st.push(5)
	fmt.Println("stack", st.elements)

	fmt.Println(st.pop())
	fmt.Println(st.pop())
	fmt.Println(st.pop())

	fmt.Println("stack", st.elements)

	q := newQueue()
	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)
	fmt.Println("queue", q.elements)

	fmt.Println("dequeue", q.dequeue())
	fmt.Println("dequeue", q.dequeue())
	fmt.Println("dequeue", q.dequeue())
	fmt.Println("q elems", q.elements)

	fmt.Println("front", q.front())
	fmt.Println("rear", q.rear())
	fmt.Println(q.elements)
}

type queue struct {
	elements []int
}

func newQueue() *queue {
	return &queue{
		elements: make([]int, 0),
	}
}

func (q *queue) enqueue(elem int) {
	q.elements = append(q.elements, elem)
}

func (q *queue) dequeue() int {
	if len(q.elements) == 0 {
		panic("queue is already empty")
	}

	elemToBeRemoved := q.elements[0]
	q.elements = q.elements[1:]

	return elemToBeRemoved
}

func (q *queue) front() int {
	if len(q.elements) == 0 {
		panic("queue is empty")
	}

	return q.elements[0]
}

func (q *queue) rear() int {

	if len(q.elements) == 0 {
		panic("queue is empty")
	}

	return q.elements[len(q.elements)-1]
}

func (q *queue) size(elem int) int {
	return len(q.elements)
}
