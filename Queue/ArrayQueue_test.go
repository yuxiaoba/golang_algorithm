package Queue

import "testing"

func TestArrayQueue_Print(t *testing.T) {
	queue := NewArrayQueue(5)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.Print()

	queue.DeQueue()
	queue.Print()

	queue.EnQueue(4)
	queue.EnQueue(5)
	queue.EnQueue(6)
	queue.EnQueue(7)
	queue.Print()
}
