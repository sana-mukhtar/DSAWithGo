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
}
