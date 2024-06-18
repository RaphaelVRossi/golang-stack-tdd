package stack_test

import (
	"testing"

	Stack "github.com/RaphaelVRossi/golang-stack-tdd/stack"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	stack := Stack.New()
	assert.NotNil(t, stack)
	assert.True(t, stack.IsEmpty())
}

func Test_Push_IsNotEmpty(t *testing.T) {
	stack := Stack.New()
	stack.Push(1)
	assert.False(t, stack.IsEmpty())
	assert.Equal(t, 1, stack.GetSize())
}

func Test_Push_Pop_isEmpty(t *testing.T) {
	stack := Stack.New()
	stack.Push(1)
	stack.Pop()
	assert.True(t, stack.IsEmpty())
	assert.Equal(t, 0, stack.GetSize())
}

func Test_Push_GetSize_IsTwo(t *testing.T) {
	stack := Stack.New()
	stack.Push(0)
	stack.Push(0)
	assert.Equal(t, 2, stack.GetSize())
}

func Test_Pop_UnderflowError(t *testing.T) {
	stack := Stack.New()
	_, err := stack.Pop()

	assert.Error(t, err)
}

func Test_Push_X_Pop_X(t *testing.T) {
	stack := Stack.New()

	stack.Push(99)
	element, _ := stack.Pop()
	assert.Equal(t, 99, element)

	stack.Push(88)
	element, _ = stack.Pop()
	assert.Equal(t, 88, element)
}

func Test_Push_XY_Pop_YX(t *testing.T) {
	stack := Stack.New()

	stack.Push(99)
	stack.Push(88)

	element, _ := stack.Pop()
	assert.Equal(t, 88, element)

	element, _ = stack.Pop()
	assert.Equal(t, 99, element)
}
