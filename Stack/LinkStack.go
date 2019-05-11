package Stack

import "fmt"

type node struct {
	next *node
	value interface{}
}

type LinkedListStack struct {
	top *node
}

func NewLinkerListStack() *LinkedListStack  {
	return  &LinkedListStack{nil}
}

func (stack *LinkedListStack) IsEmpty() bool  {
	if stack.top == nil{
		return true
	}
	return false
}

func (stack *LinkedListStack) Push(v interface{}) {
	stack.top = &node{stack.top, v}
}

func (stack *LinkedListStack) Pop() interface{}  {
	if stack.IsEmpty() {
		fmt.Println("Ths stack is empty")
		return nil
	}
	v := stack.top.value
	stack.top = stack.top.next
	return  v
}

func (stack *LinkedListStack)Flush()  {
	stack.top = nil
}

func (stack *LinkedListStack) Print()  {
	if stack.IsEmpty(){
		fmt.Println("The stack is empty")
	}else {
		current := stack.top
		for current != nil{
			fmt.Printf("%v ", current.value)
			current = current.next
		}
		fmt.Println()
	}
}