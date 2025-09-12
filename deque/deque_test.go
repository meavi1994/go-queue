package deque

import (
	"fmt"
	"testing"
)

func TestDeque(t *testing.T) {
	dq := New[int]()

	dq.Push(10)
	dq.Push(20)
	dq.PushFront(5)

	if dq.Len() != 3 {
		t.Errorf("expected length 3, got %v", dq.Len())
	}

	front, _ := dq.PopFront()
	if front != 5 {
		t.Errorf("expected front 5, got %v", front)
	}

	back, _ := dq.Pop()
	if back != 20 {
		t.Errorf("expected back 20, got %v", back)
	}

	dq.Clear()
	if !dq.IsEmpty() {
		t.Errorf("expected empty deque after clear")
	}
}

type El struct {
	ID  int
	Val int
}

func (el *El) String() string {
	return fmt.Sprintf("(%v,%v)", el.ID, el.Val)
}

func TestPriorityQueue_String(t *testing.T) {
	pq := New[*El]()
	pq.PushFront(&El{
		ID:  4,
		Val: 0,
	})
	pq.PushFront(&El{
		ID:  0,
		Val: 2,
	})
	pq.Push(&El{
		ID:  3,
		Val: 0,
	})
	pq.PushBack(&El{
		ID:  2,
		Val: 0,
	})
	fmt.Println(pq)
}
