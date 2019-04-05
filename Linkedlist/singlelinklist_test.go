package Linkedlist

import "testing"

func TestLinkedList_InsertHead(t *testing.T) {
	list := NewLinkedList()
	for i:=0; i<10; i++ {
		list.InsertHead(i+1)
	}
	list.Print()
}

func TestLinkedList_InsertTail(t *testing.T) {
	list := NewLinkedList()
	for i:=0; i<10; i++ {
		list.InsertTail(i+1)
	}
	list.Print()
}

func TestLinkedList_DeleteNode(t *testing.T) {
	list := NewLinkedList()
	for i:=0; i<3; i++{
		list.InsertTail(i+1)
	}
	list.Print()

	t.Log(list.DeleteNode(list.head.next.next))
	list.Print()
}


