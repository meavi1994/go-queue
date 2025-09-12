package priorityqueue

import (
	"container/heap"
	"fmt"
	"strings"
)

// PriorityQueue is a generic, non-thread-safe priority queue.
type PriorityQueue[T any] struct {
	items []*item[T]
	less  func(a, b T) bool
}

type item[T any] struct {
	value T
	index int // internal index
}

// New creates a new priority queue with a custom less function.
func New[T any](less func(a, b T) bool) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{less: less, items: []*item[T]{}}
	heap.Init(pq)
	return pq
}

// Len returns the number of items.
func (pq PriorityQueue[T]) Len() int { return len(pq.items) }
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq.less(pq.items[i].value, pq.items[j].value)
}
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}
func (pq *PriorityQueue[T]) Push(x any) {
	it := x.(*item[T])
	it.index = len(pq.items)
	pq.items = append(pq.items, it)
}
func (pq *PriorityQueue[T]) Pop() any {
	n := len(pq.items)
	it := pq.items[n-1]
	it.index = -1
	pq.items = pq.items[:n-1]
	return it
}

// Push adds a value to the queue.
func (pq *PriorityQueue[T]) PushValue(value T) {
	heap.Push(pq, &item[T]{value: value})
}

// Pop removes and returns the top-priority value.
func (pq *PriorityQueue[T]) PopValue() (T, bool) {
	if pq.Len() == 0 {
		var zero T
		return zero, false
	}
	it := heap.Pop(pq).(*item[T])
	return it.value, true
}

// Peek returns the top-priority value without removing it.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.Len() == 0 {
		var zero T
		return zero, false
	}
	return pq.items[0].value, true
}

// Peek returns the top-priority value without removing it.
func (pq *PriorityQueue[T]) PeekValue() (T, bool) {
	return pq.Peek()
}

// String implements fmt.Stringer
func (pq *PriorityQueue[T]) String() string {
	var sb strings.Builder
	sb.WriteString("PriorityQueue [")
	for i, it := range pq.items {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(fmt.Sprintf("%v", it.value))
	}
	sb.WriteString("]")
	return sb.String()
}
