package Stack

import (
	"testing"
	"fmt"
)

func TestArrayStack_Push(t *testing.T) {
	stack := NewArrayStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
    stack.Print()
	fmt.Println()

}

func TestArrayStack_Pop(t *testing.T) {
	stack := NewArrayStack()
	stack.Push(1)
	stack.Push(2)
	stack.Print()
	fmt.Println()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Print()
	fmt.Println()

}