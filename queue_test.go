package container

import "testing"

func TestQueueLIFO(t *testing.T) {
	q := NewQueue(LIFO)
	if q.Length() != 0 {
		t.Errorf("New queue is not empty.\n")
	}
	if !q.Empty() {
		t.Errorf("New queue is not empty.\n")
	}
	if q.Type() != LIFO {
		t.Errorf("New queue is not LIFO.\n")
	}
}

func TestQueueFIFO(t *testing.T) {
	q := NewQueue(FIFO)
	if q.Length() != 0 {
		t.Errorf("New queue is not empty.\n")
	}
	if !q.Empty() {
		t.Errorf("New queue is not empty.\n")
	}
	if q.Type() != FIFO {
		t.Errorf("New queue is not FIFO.\n")
	}
}

func TestQueueLIFOAccess(t *testing.T) {
	q := NewQueue(LIFO)
	values := []int{1, 3, 5, 7, 9, 2, 4, 6}

	for i, v := range values {
		q.Push(v)
		if q.Length() != i+1 {
			t.Errorf("Queue length is %v, expected %v.\n", q.Length(), i+1)
		}
	}

	for i, v := range values {
		value := q.Pop()
		if value != v {
			t.Errorf("Value is %v, expected %v.\n", value, i)
		}
		if q.Length() != len(values)-i-1 {
			t.Errorf("Queue length is %v, expected %v.\n", q.Length(), len(values)-i-1)
		}
	}
}

func TestQueueFIFOAccess(t *testing.T) {
	q := NewQueue(FIFO)
	values := []int{1, 3, 5, 7, 9, 2, 4, 6}

	for i, v := range values {
		q.Push(v)
		if q.Length() != i+1 {
			t.Errorf("Queue length is %v, expected %v.\n", q.Length(), i+1)
		}
	}

	for i := range values {
		value := q.Pop()
		if value != values[len(values)-1-i] {
			t.Errorf("Value is %v, expected %v.\n", value, i)
		}
		if q.Length() != len(values)-i-1 {
			t.Errorf("Queue length is %v, expected %v.\n", q.Length(), len(values)-i-1)
		}
	}
}
