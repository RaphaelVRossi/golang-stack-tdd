package golang_stack_tdd_test

import (
	"testing"

	"github.com/RaphaelVRossi/golang_stack_tdd"
	"github.com/stretchr/testify/assert"
)

func TestCreateStack(t *testing.T) {
	stack := NewStack()
	assert.NotNil(t, stack)
	assert.True(t, stack.IsEmpty())
}

func TestAfterOnePush_isNotEmpty(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	assert.False(t, stack.IsEmpty())
	assert.Equal(t, 1, stack.GetSize())
}

func TestAfterOnePushAndOnePop_isEmpty(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Pop()
	assert.True(t, stack.IsEmpty())
	assert.Equal(t, 0, stack.GetSize())
}

func TestAfterTwoPushes_sizeIsTwo(t *testing.T) {
	stack := NewStack()
	stack.Push(0)
	stack.Push(0)
	assert.Equal(t, 2, stack.GetSize())
}

func TestPoppingEmptyStack_ErrorUnderflow(t *testing.T) {
	stack := NewStack()
	_, err := stack.Pop()

	assert.Error(t, err)
}

func TestAfterPushingX_willPopX(t *testing.T) {
	stack := NewStack()

	stack.Push(99)
	element, _ := stack.pop()
	assert.Equal(t, 99, element)

	stack.Push(88)
	element, _ = stack.pop()
	assert.Equal(t, 88, element)
}

func TestAfterPushingXandY_willPopYthenX(t *testing.T) {
	stack := NewStack()

	stack.Push(99)
	stack.Push(88)

	element, _ := stack.Pop()
	assert.Equal(t, 88, element)

	element, _ = stack.Pop()
	assert.Equal(t, 99, element)
}
