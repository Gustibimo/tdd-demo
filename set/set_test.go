package set

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

//Specs:
//	- Set is empty on construction
//	- Set has 0 'size'' on construction
//	- Should perform 'add'' x value to empty set, after add operation size should be equal n
//	- Should perform 'remove'' x value and return removed value, if one add x value then remove
//	- Should return True if a set has x value, the size of set doesnt change.
//	- Should return ErrNoElement, if one perform remove to empty set.
//	- Should return ErrNoLEment, if one perform isIn to emptyset.

func TestSet_New(t *testing.T) {
	t.Run("Set is empty on construction", func(t *testing.T) {
		s := New()
		assert.True(t, s.isEmpty())
	})
	t.Run("Set has 0 'size'' on construction", func(t *testing.T) {
		s := New()
		assert.Equal(t, 0, s.Size())
	})
}

func TestSet_Operations(t *testing.T) {
	t.Run("Should perform 'add'' x value to empty set, after add operation size should be equal n", func(t *testing.T) {
		s := New()
		_ = s.Add("1")
		_ = s.Add("satu")
		_ = s.Add("bukan satu")
		assert.False(t, s.isEmpty())
		assert.Equal(t, 3, s.Size())
	})

	t.Run("Should perform 'remove' x value and return removed value, if one add x value then remove", func(t *testing.T) {
		s := New()
		_ = s.Add("1")
		_ = s.Add("satu")
		_ = s.Add("bukan satu")
		assert.Equal(t, 3, s.Size())
		value, _ := s.Remove("1")
		assert.Equal(t, "1", value)
		assert.Equal(t, 2, s.Size())
	})

	t.Run("Should return True if a set has x value, the size of set doesnt change", func(t *testing.T) {
		s := New()
		_ = s.Add("1")
		_ = s.Add("satu")
		_ = s.Add("bukan satu")
		assert.Equal(t, 3, s.Size())
		value, _ := s.Has("1")
		assert.True(t, value)
		assert.Equal(t, 3, s.Size())
	})

}
