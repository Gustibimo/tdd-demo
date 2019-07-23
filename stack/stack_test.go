package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Requirements:
//- Stack is empty on construction
//- stack has size 0 on construction
//- can perform pushes to empty stack, after push stack size equal n
//- if one push x then pop, the value popped is x and the size is one less then before(old size)
//- if one push x then peek, the value returned is x, but the size is still not change
//- if the size is n, after n pops, then the stack will be empty and has size 0
//- Popping from the empty stack should return ErrNoSuchElement
//- Peeking into the empty stack should return exception ErroSuchElement

func TestStack_NewStack(t *testing.T) {
	t.Run("Stack is empty on construction", func(t *testing.T) {
		s := New()
		assert.True(t, s.IsEmpty())
	})

	t.Run("stack has size 0 on construction", func(t *testing.T) {
		s := New()
		assert.Equal(t, 0, s.Size())
	})
}

func TestStack_Insert(t *testing.T) {
	t.Run("can perform pushes to empty stack, after push stack size equal n", func(t *testing.T) {
		s := New()
		s.Push(1)
		s.Push(2)
		s.Push(3)
		assert.False(t, s.IsEmpty())
		assert.Equal(t, 3, s.Size())

	})

	t.Run("if one push x then pop, the value popped is x and the size is one less then before(old size)", func(t *testing.T) {
		s := New()
		s.Push(1)
		s.Push(2)
		s.Push(4)
		assert.Equal(t, 3, s.Size())
		value, _ := s.Pop()
		assert.Equal(t, 4, value)
		assert.Equal(t, 2, s.Size())
	})

	t.Run("if one push x then peek, the value returned is x, but the size is still not change", func(t *testing.T) {
		s := New()
		s.Push(1)
		s.Push(2)
		s.Push(4)
		assert.Equal(t, 3, s.Size())
		value, _ := s.Peek()
		assert.Equal(t, 4, value)
		assert.Equal(t, 3, s.Size())
	})

}

func TestStack_Error(t *testing.T) {

	t.Run("Popping from the empty stack should return ErrNoSuchElement", func(t *testing.T) {
		s := New()
		_, err := s.Pop()
		if err == nil {
			t.Fail()
			t.Logf("expect error is not nil, but got `%v`", err)
		}
		assert.Equal(t, ErrNotSuchElement, err)
	})

	t.Run("Peeking into the empty stack should return exception ErroSuchElement", func(t *testing.T) {
		s := New()
		_, err := s.Peek()
		if err == nil {
			t.Fail()
			t.Logf("expect error is not nil, but got `%v`", err)
		}
		assert.Equal(t, ErrNotSuchElement, err)
	})

}
