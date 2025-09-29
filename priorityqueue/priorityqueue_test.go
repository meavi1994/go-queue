package priorityqueue

import (
	"fmt"
	"testing"
)

func TestPriorityQueue(t *testing.T) {
	pq := New[int](func(a, b int) bool { return a > b }) // max-heap

	pq.PushValue(10)
	pq.PushValue(5)
	pq.PushValue(20)
	value := pq.PushAndReturnItem(30)
	pq.RemoveItem(value)
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

type El struct {
	ID  int
	Val int
}

func (el *El) String() string {
	return fmt.Sprintf("(%v,%v)", el.ID, el.Val)
}

func TestPriorityQueue_String(t *testing.T) {
	pq := New[*El](func(a, b *El) bool { return a.ID > b.ID })
	pq.PushValue(&El{
		ID:  4,
		Val: 0,
	})
	pq.PushValue(&El{
		ID:  0,
		Val: 2,
	})
	pq.PushValue(&El{
		ID:  3,
		Val: 0,
	})
	pq.PushValue(&El{
		ID:  2,
		Val: 0,
	})
	fmt.Println(pq)
}
