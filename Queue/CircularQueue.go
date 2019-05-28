package Queue

import "fmt"

// 循环队列

type CircularQueue struct {
	queue []interface{}
	head int
	tail int
	cap int
}

func NewCircularQueue(n int) *CircularQueue {
	return &CircularQueue{
		make([]interface{}, n ), 0,0,n,
	}
}

func (queue *CircularQueue)EnQueue(v interface{}) bool {
	if (queue.tail+1)%queue.cap == queue.head{
		fmt.Println("CircularQueue is full")
		return false
	}
	queue.queue[queue.tail] = v
	queue.tail = (queue.tail+1)%queue.cap
	return true
}

func (queue *CircularQueue)DeQueue() interface{} {
	if queue.tail == queue.head{
		fmt.Println("CircularQueue is empty")
		return nil
	}
	v := queue.queue[queue.head]
	queue.head = (queue.head+1)%queue.cap
	return v
}

func (queue *CircularQueue) Print() {
	if queue.tail == queue.head{
		fmt.Println("CircularQueue is empty")
	}
	for i := queue.head; ;{
		fmt.Printf("%+v", queue.queue[i])
		i = (i+1)%queue.cap
		if i==queue.tail{
			break
		}
	}
	fmt.Println()
}