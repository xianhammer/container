package container

type QueueType int

type Queue struct {
	kind        QueueType
	length      int
	first, last *Node
}

const (
	LIFO QueueType = 1
	FIFO           = 2
)

// NewQueue create a new queue
func NewQueue(kind QueueType) (q *Queue) {
	q = new(Queue)
	q.kind = kind
	return
}

// Length of queue
func (q *Queue) Type() QueueType {
	return q.kind
}

// Length of queue
func (q *Queue) Length() int {
	return q.length
}

// Empty return true if queue holds no elements
func (q *Queue) Empty() bool {
	return q.length == 0
}

// Push a value into the queue
func (q *Queue) Push(value interface{}) {
	n := NewNode(value)
	if q.last == nil {
		q.last = n
	}
	if q.first != nil {
		q.first.InsertAfter(n)
	}
	q.first = n
	q.length++
}

// Pop a value from the queue. No boundary checks are performed!
func (q *Queue) Pop() (value interface{}) {
	switch q.kind {
	case LIFO:
		value = q.last.Value
		q.last = q.last.prev
	case FIFO:
		value = q.first.Value
		q.first = q.first.next
	}
	q.length--
	return
}
