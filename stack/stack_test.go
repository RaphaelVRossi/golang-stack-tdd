package stack_test

import (
	"testing"

	"github.com/RaphaelVRossi/golang-stack-tdd/stack"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := stack.New()
	assert.NotNil(t, s)
	assert.True(t, s.IsEmpty())
}

func Test_Push_IsNotEmpty(t *testing.T) {
	s := stack.New()
	s.Push(1)
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 1, s.GetSize())
}

func Test_Push_Pop_isEmpty(t *testing.T) {
	s := stack.New()
	s.Push(1)
	s.Pop()
	assert.True(t, s.IsEmpty())
	assert.Equal(t, 0, s.GetSize())
}

func Test_Push_GetSize_IsTwo(t *testing.T) {
	s := stack.New()
	s.Push(0)
	s.Push(0)
	assert.Equal(t, 2, s.GetSize())
}

func Test_Pop_UnderflowError(t *testing.T) {
	s := stack.New()
	_, err := s.Pop()

	assert.Error(t, err)
}

func Test_Push_X_Pop_X(t *testing.T) {
	s := stack.New()

	s.Push(99)
	element, _ := s.Pop()
	assert.Equal(t, 99, element)

	s.Push(88)
	element, _ = s.Pop()
	assert.Equal(t, 88, element)
}

func Test_Push_XY_Pop_YX(t *testing.T) {
	s := stack.New()

	s.Push(99)
	s.Push(88)

	element, _ := s.Pop()
	assert.Equal(t, 88, element)

	element, _ = s.Pop()
	assert.Equal(t, 99, element)
}
