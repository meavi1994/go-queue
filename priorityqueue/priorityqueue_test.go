package priorityqueue

import (
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := New[int](func(a, b int) bool { return a > b }) // max-heap

	pq.PushValue(10)
	pq.PushValue(5)
	pq.PushValue(20)

	if val, _ := pq.Peek(); val != 20 {
		t.Errorf("expected 20, got %v", val)
	}

	expected := []int{20, 10, 5}
	for _, v := range expected {
		val, ok := pq.PopValue()
		if !ok || val != v {
			t.Errorf("expected %v, got %v", v, val)
		}
	}

	if pq.Len() != 0 {
		t.Errorf("expected length 0, got %v", pq.Len())
	}
}
