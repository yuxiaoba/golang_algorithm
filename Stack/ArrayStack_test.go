package Stack

import (
	"testing"
)

func TestArrayStack_Push(t *testing.T) {
	stack := NewArrayStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
    stack.Print()

}

func TestArrayStack_Pop(t *testing.T) {
	stack := NewArrayStack()
	stack.Push(1)
	stack.Push(2)
	stack.Print()

	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Print()


}