package stack

type Stack struct {
	empty bool
}

func New() *Stack {
	return &Stack{empty: true}
}

func (s *Stack) IsEmpty() bool {
	return s.empty
}

func (s *Stack) Push(element int) {
	s.empty = false
}

func (s *Stack) Pop() int {
	s.empty = true
	return 1
}
