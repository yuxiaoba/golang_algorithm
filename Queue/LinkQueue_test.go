package Queue

import (
	"testing"
	"fmt"
)

func TestLinkedListQueue_Print(t *testing.T) {
	queue := NewLinkerListQueue()
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.Print()
	fmt.Println(queue.len)

	queue.DeQueue()
	queue.Print()
	queue.DeQueue()
	queue.Print()
}
