package Linkedlist

import "testing"

func TestLinkedList_Reverse(t *testing.T) {
	value:= []int{1,2, 3, 4, 5, 6}

	list := NewLinkedList()
	for _, v := range value{
		list.InsertTail(v)
	}
	list.Print()

	relist := list.Reverse()
    relist.Print()
}

func TestLinkedList_IsHasCycle(t *testing.T) {
	value:= []int{1,2, 3}

	list := NewLinkedList()
	for _, v := range value{
		list.InsertTail(v)
	}
	list.Print()

	t.Log(list.IsHasCycle())
	list.head.next.next.next.next = list.head.next
	t.Log(list.IsHasCycle())
}

func TestMergeSortedList(t *testing.T) {
	value1:= []int{1, 3,  5}

	list1 := NewLinkedList()
	for _, v := range value1{
		list1.InsertTail(v)
	}
	list1.Print()

	value2 := []int{2,4, 6}
	list2 := NewLinkedList()
	for _, v := range value2{
		list2.InsertTail(v)
	}
	list2.Print()

	MergeSortedList(list1,list2).Print()
}


func TestLinkedList_DeleteBottleN(t *testing.T) {
	value:= []int{1, 2, 3, 4, 5, 6}

	list := NewLinkedList()
	for _, v := range value{
		list.InsertTail(v)
	}
	list.Print()
    list.DeleteBottleN(2)
	list.Print()
}