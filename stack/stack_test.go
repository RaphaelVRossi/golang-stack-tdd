package stack_test

import (
	"testing"

	"github.com/RaphaelVRossi/golang-stack-tdd/stack"
	"github.com/stretchr/testify/assert"
)

func TestNothing(t *testing.T) {
	assert.True(t, 1 == 1)
}

func TestNew(t *testing.T) {
	s := stack.New()
	assert.True(t, s.IsEmpty())
}

func TestPushIsNotEmpty(t *testing.T) {
	s := stack.New()

	s.Push(1)
	assert.False(t, s.IsEmpty())
}

func TestPushPopIsEmpty(t *testing.T) {
	s := stack.New()

	s.Push(1)
	e := s.Pop()

	assert.True(t, s.IsEmpty())
	assert.Equal(t, 1, e)
}
