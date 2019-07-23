package stack

import "fmt"

var ErrNotSuchElement = fmt.Errorf("Not such element")

func New() *Stack {

	return &Stack{
		arr: make([]int, 0),
	}

}

type Stack struct {
	arr []int
}

func (s Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s Stack) Size() int {
	return len(s.arr)
}

func (s *Stack) Push(value int) error {
	s.arr = append(s.arr, value)
	return nil
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrNotSuchElement
	}
	lastIndex := s.getLastIndex()
	s.arr = s.arr[:len(s.arr)-1]
	return lastIndex, nil
}

func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, ErrNotSuchElement
	}
	value := s.getLastIndex()
	return value, nil
}

func (s *Stack) getLastIndex() int {
	value := s.arr[len(s.arr)-1]
	return value
}
