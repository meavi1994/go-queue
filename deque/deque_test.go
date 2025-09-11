package deque

import (
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
