package priorityqueue

import (
	"container/heap"
	"fmt"
	"strings"
)

// PriorityQueue is a generic, non-thread-safe priority queue.
type PriorityQueue[T any] struct {
	items []*Item[T]
	less  func(a, b T) bool
}

type Item[T any] struct {
	Value T
	index int // internal index
}

// New creates a new priority queue with a custom less function.
func New[T any](less func(a, b T) bool) *PriorityQueue[T] {
	pq := &PriorityQueue[T]{less: less, items: []*Item[T]{}}
	heap.Init(pq)
	return pq
}

// Len returns the number of items.
func (pq PriorityQueue[T]) Len() int { return len(pq.items) }
func (pq PriorityQueue[T]) Less(i, j int) bool {
	return pq.less(pq.items[i].Value, pq.items[j].Value)
}
func (pq PriorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}
func (pq *PriorityQueue[T]) Push(x any) {
	it := x.(*Item[T])
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

// Push adds a Value to the queue.
func (pq *PriorityQueue[T]) PushValue(value T) {
	pq.PushAndReturnItem(value)
}

func (pq *PriorityQueue[T]) PushAndReturnItem(value T) *Item[T] {
	it := &Item[T]{Value: value}
	heap.Push(pq, it)
	return it
}

func (pq *PriorityQueue[T]) RemoveItem(it *Item[T]) (T, bool) {
	if it.index < 0 || it.index >= pq.Len() {
		var zero T
		return zero, false
	}
	removed := heap.Remove(pq, it.index).(*Item[T])
	return removed.Value, true
}

// Pop removes and returns the top-priority Value.
func (pq *PriorityQueue[T]) PopValue() (T, bool) {
	if pq.Len() == 0 {
		var zero T
		return zero, false
	}
	it := heap.Pop(pq).(*Item[T])
	return it.Value, true
}

// Peek returns the top-priority Value without removing it.
func (pq *PriorityQueue[T]) Peek() (T, bool) {
	if pq.Len() == 0 {
		var zero T
		return zero, false
	}
	return pq.items[0].Value, true
}

// Peek returns the top-priority Value without removing it.
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
		sb.WriteString(fmt.Sprintf("%v", it.Value))
	}
	sb.WriteString("]")
	return sb.String()
}
