package tdd_stack

import (

	"github.com/gustibimo/tdd-stack/stack"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Requirements:
//- Stack is empty on construction
//- tdd-stack has size 0 on construction
//- can perform pushes to empty tdd-stack, after push tdd-stack size equal n
//- if one push x then pop, the value popped is x and the size is one less then before(old size)
//- if one push x then peek, the value returned is x, but the size is still not change
//- if the size is n, after n pops, then the tdd-stack will be empty and has size 0
//- Popping from the empty tdd-stack should return ErrNoSuchElement
//- Peeking into the empty tdd-stack should return exception ErroSuchElement

func TestNewSet(t *testing.T) {
	t.Run("Stack is empty on construction", func(t *testing.T) {
		s := stack.New()
		assert.True(t, s.IsEmpty())
	})
}