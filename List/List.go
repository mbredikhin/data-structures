package List

import (
	"fmt"
)

// List data structure
type List struct {
	memory []int
	length int
}

// NewList - List constructor
func NewList(memory []int, length int) *List {
	return &List{memory, length}
}

// Get value by address
func (l *List) Get(address int) (int, error) {
	if l.length-1 < address || address < 0 {
		return 0, fmt.Errorf("Can't get value by address %d", address)
	}
	return l.memory[address], nil
}

// Push value to the end of list
func (l *List) Push(value int) {
	l.memory = append(l.memory, value)
	l.length++
}

// Pop - remove value from the end of list and return it
func (l *List) Pop() (int, error) {
	if l.length == 0 {
		return 0, fmt.Errorf("Can't pop value from empty list")
	}
	return l.Get(l.length - 1)
}

// Unshift - insert value to the top of list
func (l *List) Unshift(value int) {
	l.memory = append(l.memory, value)
	l.length++
}

// Shift - remove value from the top of list and return it
func (l *List) Shift() (int, error) {
	if l.length == 0 {
		return 0, fmt.Errorf("Can't unshift value from empty list")
	}
	el, err := l.Get(0)
	if err != nil {
		return 0, err
	}
	l.memory = l.memory[1:]
	l.length--
	return el, nil
}
