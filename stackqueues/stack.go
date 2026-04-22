package stackqueues

type Stack struct {
	elements []interface{}
}

func NewStack() *Stack {
	return &Stack{
		elements: make([]interface{}, 0),
	}
}

func (st *Stack) Push(elem interface{}) {
	st.elements = append(st.elements, elem)
}

func (st *Stack) Pop(elem interface{}) interface{} {
	if len(st.elements) == 0 {
		return nil
	}

	lastIdx := len(st.elements) - 1
	lastElem := st.elements[lastIdx]

	st.elements[lastIdx] = nil // this is optional, but it's safer to implement to avoid memory leaks
	st.elements = st.elements[:lastIdx]

	return lastElem
}

func (st *Stack) Peek() interface{} {
	if len(st.elements) == 0 {
		return nil
	}

	lastIdx := len(st.elements) - 1
	lastElem := st.elements[lastIdx]

	return lastElem
}

func (st *Stack) IsEmpty() bool {
	return len(st.elements) == 0
}

func (st *Stack) Size() int {
	return len(st.elements)
}
