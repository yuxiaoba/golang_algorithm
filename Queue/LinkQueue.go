package Queue

import "fmt"

// 链表队列

type Node struct {
	v interface{}
	next *Node
}

type LinkedListQueue struct {
	head *Node
	tail *Node
	len int
}

func NewLinkerListQueue()*LinkedListQueue  {
	return &LinkedListQueue{nil, nil,0}
}

func (Q *LinkedListQueue)EnQueue (v interface{}) {
	node := &Node{v, nil}
	if Q.len ==0 {
		// 队列为空
		Q.head = node
		Q.tail = node
	}else {
		Q.tail.next = node
		Q.tail = node
	}
	Q.len ++
}

func (Q *LinkedListQueue)DeQueue() interface{} {
	if Q.len == 0 {
		fmt.Println("Queue is empty")
		return nil
	}
	v := Q.head.v
	Q.head = Q.head.next
	Q.len --
	return v
}

func (Q *LinkedListQueue)Print()  {
	if Q.len == 0 {
		fmt.Println("Queue is empty")
	}
	for cur := Q.head; cur != nil ; cur = cur.next{
		fmt.Printf("%+v ", cur.v)
	}
	fmt.Println()
}
