package examples

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStack(t *testing.T) {
	stack := Stack()
	assert.NotNil(t, stack)
	assert.True(t, stack.IsEmpty())
}

func TestAfterOnePush_isNotEmpty(t *testing.T) {
	stack := Stack()
	stack.push(1)
	assert.False(t, stack.IsEmpty())
	assert.Equal(t, 1, stack.GetSize())
}

func TestAfterOnePushAndOnePop_isEmpty(t *testing.T) {
	stack := Stack()
	stack.push(1)
	stack.pop()
	assert.True(t, stack.IsEmpty())
	assert.Equal(t, 0, stack.GetSize())
}

func TestAfterTwoPushes_sizeIsTwo(t *testing.T) {
	stack := Stack()
	stack.push(0)
	stack.push(0)
	assert.Equal(t, 2, stack.GetSize())
}

func TestPoppingEmptyStack_ErrorUnderflow(t *testing.T) {
	stack := Stack()
	_, err := stack.pop()

	assert.Error(t, err)
}

func TestAfterPushingX_willPopX(t *testing.T) {
	stack := Stack()

	stack.push(99)
	element, _ := stack.pop()
	assert.Equal(t, 99, element)

	stack.push(88)
	element, _ = stack.pop()
	assert.Equal(t, 88, element)
}

func TestAfterPushingXandY_willPopYthenX(t *testing.T) {
	stack := Stack()

	stack.push(99)
	stack.push(88)

	element, _ := stack.pop()
	assert.Equal(t, 88, element)

	element, _ = stack.pop()
	assert.Equal(t, 99, element)
}
