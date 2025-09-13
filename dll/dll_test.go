package dll

import (
	"fmt"
	"testing"
)

func TestPushAndToSlice(t *testing.T) {
	l := New[int]()
	l.PushBack(1)
	l.PushBack(2)
	l.PushFront(0)

	expected := []int{0, 1, 2}
	if got := l.ToSlice(); fmt.Sprint(got) != fmt.Sprint(expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestReverse(t *testing.T) {
	l := New[int]()
	l.FromSlice([]int{1, 2, 3})
	l.Reverse()

	expected := []int{3, 2, 1}
	if got := l.ToSlice(); fmt.Sprint(got) != fmt.Sprint(expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestSortFunc(t *testing.T) {
	l := New[int]()
	l.FromSlice([]int{3, 1, 2})
	l.SortFunc(func(a, b int) bool { return a < b })

	expected := []int{1, 2, 3}
	if got := l.ToSlice(); fmt.Sprint(got) != fmt.Sprint(expected) {
		t.Errorf("expected %v, got %v", expected, got)
	}
}

func TestStringer(t *testing.T) {
	l := New[int]()
	l.FromSlice([]int{1, 2, 3})
	if got := l.String(); got != "[1 2 3]" {
		t.Errorf("expected [1 2 3], got %v", got)
	}
}
