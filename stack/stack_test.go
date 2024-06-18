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

func TestPushIsNotEmpty(t *testing.T) {
	s := stack.New()
	s.Push(1)
	assert.False(t, s.IsEmpty())
	assert.Equal(t, 1, s.GetSize())
}

func TestPushPopisEmpty(t *testing.T) {
	s := stack.New()
	s.Push(1)
	s.Pop()
	assert.True(t, s.IsEmpty())
	assert.Equal(t, 0, s.GetSize())
}

func TestPushGetSizeIsTwo(t *testing.T) {
	s := stack.New()
	s.Push(0)
	s.Push(0)
	assert.Equal(t, 2, s.GetSize())
}

func TestPopUnderflowError(t *testing.T) {
	s := stack.New()
	_, err := s.Pop()

	assert.Error(t, err)
}

func TestPushXPopX(t *testing.T) {
	s := stack.New()

	s.Push(99)
	element, _ := s.Pop()
	assert.Equal(t, 99, element)

	s.Push(88)
	element, _ = s.Pop()
	assert.Equal(t, 88, element)
}

func TestPushXYPopYX(t *testing.T) {
	s := stack.New()

	s.Push(99)
	s.Push(88)

	element, _ := s.Pop()
	assert.Equal(t, 88, element)

	element, _ = s.Pop()
	assert.Equal(t, 99, element)
}
