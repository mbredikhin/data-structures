package Queue

import "errors"

// Queue data structure
type Queue struct {
	list   []int
	length int
}

// NewQueue - Queue constructor
func NewQueue() *Queue {
	var length int = 0
	var list = make([]int, length)
	return &Queue{list, length}
}

// Enqueue - append value to the end of queue
func (q *Queue) Enqueue(v int) {
	q.length++
	q.list = append(q.list, v)
}

// Dequeue - append value to the front of queue
func (q *Queue) Dequeue(v int) {
	q.length++
	q.list = append([]int{v}, q.list...)
}

// Peek - get element from the front of the queue
func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("There is no values in queue")
	}
	return q.list[0], nil
}

// IsEmpty - check queue for values
func (q *Queue) IsEmpty() bool {
	if q.length == 0 {
		return true
	}
	return false
}

// Empty - empty queue
func (q *Queue) Empty() {
	q.length = 0
	q.list = make([]int, q.length)
}
