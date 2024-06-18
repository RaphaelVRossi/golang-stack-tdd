package golang_stack_tdd

type Stack struct {
	empty    bool
	size     int
	elements []int
}

type UnderflowError struct{}

func (u *UnderflowError) Error() string {
	return "can not pop a empty stack"
}

func NewStack() *Stack {
	return &Stack{size: 0, elements: make([]int, 10)}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) GetSize() int {
	return s.size
}

func (s *Stack) Push(element int) {
	s.elements[s.size] = element
	s.size++
}

func (s *Stack) Pop() (int, error) {
	if s.size == 0 {
		return -1, &UnderflowError{}
	}

	s.size--
	return s.elements[s.size], nil
}
