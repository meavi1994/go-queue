// Package dll provides a simple, generic doubly-linked list implementation.
//
// Features:
//   - Generic: works with any type `T` (Go 1.18+)
//   - Typical list operations: PushFront, PushBack, InsertBefore, InsertAfter,
//     Remove, Front, Back, Len, Clear, ToSlice, String, SortFunc
//   - Iteration via node.Next / node.Prev
//
// Note: This implementation is NOT safe for concurrent use. Protect with a mutex
// if you need concurrent access.

package dll

import (
	"fmt"
	"sort"
	"strings"
)

// Node is an element in the doubly-linked list.
type Node[T any] struct {
	Value T
	prev  *Node[T]
	next  *Node[T]
}

// Prev returns the previous node (or nil).
func (n *Node[T]) Prev() *Node[T] { return n.prev }

// Next returns the next node (or nil).
func (n *Node[T]) Next() *Node[T] { return n.next }

// List is a generic doubly-linked list.
type List[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
}

// New returns an initialized empty list.
func New[T any]() *List[T] { return &List[T]{} }

// Len returns the number of elements in the list.
func (l *List[T]) Len() int { return l.len }

// Front returns the first node or nil.
func (l *List[T]) Front() *Node[T] { return l.head }

// Back returns the last node or nil.
func (l *List[T]) Back() *Node[T] { return l.tail }

// PushFront inserts v at the front and returns the new node.
func (l *List[T]) PushFront(v T) *Node[T] {
	n := &Node[T]{Value: v}
	if l.head == nil {
		l.head, l.tail = n, n
	} else {
		n.next = l.head
		l.head.prev = n
		l.head = n
	}
	l.len++
	return n
}

// PushBack inserts v at the back and returns the new node.
func (l *List[T]) PushBack(v T) *Node[T] {
	n := &Node[T]{Value: v}
	if l.tail == nil {
		l.head, l.tail = n, n
	} else {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
	}
	l.len++
	return n
}

// InsertAfter inserts v after node n and returns the inserted node.
// If n is nil, it behaves like PushBack.
func (l *List[T]) InsertAfter(n *Node[T], v T) *Node[T] {
	if n == nil {
		return l.PushBack(v)
	}
	if n == l.tail {
		return l.PushBack(v)
	}
	newNode := &Node[T]{Value: v}
	next := n.next
	newNode.prev = n
	newNode.next = next
	n.next = newNode
	if next != nil {
		next.prev = newNode
	}
	l.len++
	return newNode
}

// InsertBefore inserts v before node n and returns the inserted node.
// If n is nil, it behaves like PushFront.
func (l *List[T]) InsertBefore(n *Node[T], v T) *Node[T] {
	if n == nil {
		return l.PushFront(v)
	}
	if n == l.head {
		return l.PushFront(v)
	}
	newNode := &Node[T]{Value: v}
	prev := n.prev
	newNode.next = n
	newNode.prev = prev
	n.prev = newNode
	if prev != nil {
		prev.next = newNode
	}
	l.len++
	return newNode
}

// Remove removes node n from the list and returns its value.
// If n is nil or the node does not belong to this list, Remove does nothing and
// returns the zero value of T.
func (l *List[T]) Remove(n *Node[T]) (zero T) {
	if n == nil || l.len == 0 {
		return zero
	}
	// Disconnect neighbors
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		// n was head
		l.head = n.next
	}
	if n.next != nil {
		n.next.prev = n.prev
	} else {
		// n was tail
		l.tail = n.prev
	}
	// Help GC
	n.prev = nil
	n.next = nil
	l.len--
	return n.Value
}

// MoveToFront moves node n to the front. If n is already at front or nil, it's a no-op.
func (l *List[T]) MoveToFront(n *Node[T]) {
	if n == nil || n == l.head || l.len < 2 {
		return
	}
	l.Remove(n)
	val := n.Value
	l.PushFront(val)
}

// MoveToBack moves node n to the back. If n is already at back or nil, it's a no-op.
func (l *List[T]) MoveToBack(n *Node[T]) {
	if n == nil || n == l.tail || l.len < 2 {
		return
	}
	l.Remove(n)
	val := n.Value
	l.PushBack(val)
}

// ToSlice returns a slice with the list elements in order.
func (l *List[T]) ToSlice() []T {
	out := make([]T, 0, l.len)
	for e := l.head; e != nil; e = e.next {
		out = append(out, e.Value)
	}
	return out
}

// FromSlice replaces the list contents with elements from s.
func (l *List[T]) FromSlice(s []T) {
	l.Clear()
	for _, v := range s {
		l.PushBack(v)
	}
}

// Clear removes all elements from the list.
func (l *List[T]) Clear() {
	for e := l.head; e != nil; {
		n := e.next
		e.prev = nil
		e.next = nil
		e = n
	}
	l.head = nil
	l.tail = nil
	l.len = 0
}

// Find finds the first node that satisfies predicate f and returns it (or nil).
func (l *List[T]) Find(f func(T) bool) *Node[T] {
	for e := l.head; e != nil; e = e.next {
		if f(e.Value) {
			return e
		}
	}
	return nil
}

// Reverse reverses the list in-place.
func (l *List[T]) Reverse() {
	if l.len < 2 {
		return
	}
	cur := l.head
	for cur != nil {
		cur.prev, cur.next = cur.next, cur.prev
		cur = cur.prev // because we swapped
	}
	l.head, l.tail = l.tail, l.head
}

// String returns a string representation of the list values.
func (l *List[T]) String() string {
	var sb strings.Builder
	sb.WriteString("[")
	for e := l.head; e != nil; e = e.next {
		sb.WriteString(fmt.Sprintf("%v", e.Value))
		if e.next != nil {
			sb.WriteString(" ")
		}
	}
	sb.WriteString("]")
	return sb.String()
}

// SortFunc sorts the list in place using the provided less function.
// less(a, b) should return true if a < b.
func (l *List[T]) SortFunc(less func(a, b T) bool) {
	if l.len < 2 {
		return
	}
	// Convert to slice
	slice := l.ToSlice()
	sort.Slice(slice, func(i, j int) bool { return less(slice[i], slice[j]) })
	// Rebuild list
	l.FromSlice(slice)
}

// Example usage (not executed):
//
//  l := dll.New[int]()
//  l.PushBack(3)
//  l.PushBack(1)
//  l.PushBack(2)
//  fmt.Println(l) // [3 1 2]
//  l.SortFunc(func(a, b int) bool { return a < b })
//  fmt.Println(l) // [1 2 3]
