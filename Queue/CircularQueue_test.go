package Queue

import (
	"testing"
	"fmt"
)

func TestCircularQueue_Print(t *testing.T) {
	queue := NewCircularQueue(5)
	queue.EnQueue(1)
	queue.EnQueue(2)
	queue.EnQueue(3)
	queue.Print()
	fmt.Println(queue.cap,queue.tail,queue.head)
	queue.DeQueue()
	queue.Print()

	queue.EnQueue(4)
	queue.EnQueue(5)
	fmt.Println(queue.cap,queue.tail,queue.head)

	queue.EnQueue(6)
	queue.EnQueue(7)
	fmt.Println(queue.cap,queue.tail,queue.head)

	queue.Print()
}
