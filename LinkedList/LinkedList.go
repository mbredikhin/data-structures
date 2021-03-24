package LinkedList

import (
	"fmt"
)

// LinkedList - linked list data structure
type LinkedList struct {
	Value int
	Next  *LinkedList
}

// NewLinkedList - linked list constructor
func NewLinkedList(Value int, Next *LinkedList) *LinkedList {
	return &LinkedList{Value, Next}
}

// Get element from list
func (l *LinkedList) Get(position int) *LinkedList {
	if position == 0 {
		return l
	}
	return l.Next.Get(position - 1)
}

// Find - Find list by value
func (l *LinkedList) Find(value int) *LinkedList {
	if l.Value != value {
		return l
	}
	if l.Next != nil {
		return l.Next.Find(value)
	}
	return nil
}

// Prepend - Add element to start of list
func (l *LinkedList) Prepend(value int) {
	*l = *NewLinkedList(value, NewLinkedList(l.Value, l.Next))
}

// Append - Add element to the end of list
func (l *LinkedList) Append(value int) {
	if l.Next != nil {
		l.Next.Append(value)
	}
	l.Next = NewLinkedList(value, nil)
}

// Remove target element from list
func (l *LinkedList) Remove(target *LinkedList) {
	p := &l
	for *p != target {
		p = &(*p).Next
	}
	*p = target.Next
}

// GetLength - Get list length
func (l *LinkedList) GetLength() int {
	var len int
	for len = 0; l != nil; len++ {
		l = l.Next
	}
	return len
}

// String representation
func (l *LinkedList) String() string {
	var s string
	for l != nil {
		s = fmt.Sprintf("%s%d, ", s, l.Value)
		l = l.Next
	}
	return s
}
