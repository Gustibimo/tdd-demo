package tdd_stack

func New() *Stack {

	return &Stack{
		arr: make([]int, 0),
	}

}

type Stack struct {
	arr []int
}

func (s Stack) IsEmpty()  bool {
	return false
}


