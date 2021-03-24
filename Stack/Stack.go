package Stack

import "errors"

// Stack data structure
type Stack struct {
	list   []int
	length int
}

// NewStack - Stack constructor
func NewStack() *Stack {
	var length int = 0
	var list = make([]int, length)
	return &Stack{list, length}
}

// Pop - remove value from the top of stack and return it
func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("There is no values in stack")
	}
	s.length--
	v := s.list[s.length]
	s.list = s.list[:s.length]
	return v, nil
}

// Push - push value to the stack
func (s *Stack) Push(v int) {
	s.list = append(s.list, v)
	s.length++
}

// Peek - get element from the top of the stack
func (s *Stack) Peek() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("There is no values in stack")
	}
	return s.list[s.length-1], nil
}

// IsEmpty - check stack for values
func (s *Stack) IsEmpty() bool {
	if s.length == 0 {
		return true
	}
	return false
}

// Empty - empty stack
func (s *Stack) Empty() {
	s.length = 0
	s.list = make([]int, s.length)
}
