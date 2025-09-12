package deque

import (
	"fmt"
	"strings"
)

// Deque is a generic, non-thread-safe double-ended queue.
type Deque[T any] struct {
	items []T
}

// New creates a new empty deque.
func New[T any]() *Deque[T] {
	return &Deque[T]{items: make([]T, 0)}
}

func (d *Deque[T]) Push(item T) {
	d.PushBack(item)
}

func (d *Deque[T]) Pop() (T, bool) {
	return d.PopBack()
}

func (d *Deque[T]) Peek() (T, bool) {
	return d.PeekBack()
}

// PushBack adds an element to the back.
func (d *Deque[T]) PushBack(item T) {
	d.items = append(d.items, item)
}

// PopBack removes and returns the element at the back.
func (d *Deque[T]) PopBack() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	last := len(d.items) - 1
	item := d.items[last]
	d.items = d.items[:last]
	return item, true
}

// PushFront adds an element to the front.
func (d *Deque[T]) PushFront(item T) {
	d.items = append([]T{item}, d.items...)
}

// PopFront removes and returns the element at the front.
func (d *Deque[T]) PopFront() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item, true
}

// PeekFront returns the front element without removing.
func (d *Deque[T]) PeekFront() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	return d.items[0], true
}

// PeekBack returns the back element without removing.
func (d *Deque[T]) PeekBack() (T, bool) {
	if len(d.items) == 0 {
		var zero T
		return zero, false
	}
	return d.items[len(d.items)-1], true
}

// Len returns the number of elements.
func (d *Deque[T]) Len() int {
	return len(d.items)
}

// IsEmpty returns true if empty.
func (d *Deque[T]) IsEmpty() bool {
	return len(d.items) == 0
}

// Clear removes all elements.
func (d *Deque[T]) Clear() {
	d.items = make([]T, 0)
}

// String implements fmt.Stringer
func (d *Deque[T]) String() string {
	var sb strings.Builder
	sb.WriteString("Deque [")
	for i, v := range d.items {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", v))
	}
	sb.WriteString("]")
	return sb.String()
}
