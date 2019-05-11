package Stack

import (
	"testing"
)

func TestLinkedListStack_Push(t *testing.T) {
	stack := NewLinkerListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()
}


func TestLinkedListStack_Pop(t *testing.T) {
	stack := NewLinkerListStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Print()

	stack.Pop()
	stack.Pop()
	stack.Print()
	stack.Pop()
	stack.Pop()
	stack.Print()

}