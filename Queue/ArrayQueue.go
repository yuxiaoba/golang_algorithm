package Queue

import "fmt"

type ArrayQueue struct {
	queue []interface{}
	cap int
	head int  // 队头下标
	tail int  // 队尾下标
}

func NewArrayQueue(n int) *ArrayQueue  {
	return &ArrayQueue{
		make([]interface{},n),
		n,
		0,
		0,
	}
}

func (queue *ArrayQueue)EnQueue(v interface{}) bool {
	if queue.tail == queue.cap{
		// 避免每次出队都移动数据
		if queue.head == 0{
			// 队满
			fmt.Println("Queue is full")
			return false
		}
		for i:=queue.head; i < queue.tail; i++{
			queue.queue[i-queue.head] = queue.queue[i]
		}
		queue.tail -= queue.head
		queue.head = 0
	}
	queue.queue[queue.tail] = v
	queue.tail++
	return true
}

func (queue *ArrayQueue)DeQueue() interface{} {
	if queue.head == queue.tail{
		// 队空
		return nil
	}
	v := queue.head
	queue.head++
	return v
}

func (queue *ArrayQueue)Print()  {
	if queue.head == queue.tail{
		fmt.Println("Queue is empty")
	}
	for i:= queue.head; i< queue.tail; i++{
		fmt.Printf("%+v ",queue.queue[i])
	}
	fmt.Println()
}