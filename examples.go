package examples

type stack struct {
	empty    bool
	size     int
	elements []int
}

type UnderflowError struct {
	msg string
}

func Stack() *stack {
	return &stack{empty: true, size: 0, elements: make([]int, 10)}
}

func (s *stack) IsEmpty() bool {
	return s.empty
}

func (s *stack) GetSize() int {
	return s.size
}

func (s *stack) push(element int) {
	s.empty = false
	s.elements[s.size] = element
	s.size++
}

func (s *stack) pop() (int, error) {
	if s.size == 0 {
		return -1, &UnderflowError{msg: "Underflow"}
	}
	s.empty = true
	s.size--
	return s.elements[s.size], nil
}

func (e *UnderflowError) Error() string {
	return e.msg
}
