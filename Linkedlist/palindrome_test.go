package Linkedlist

import "testing"

func TestIsPalindrome1(t *testing.T) {
	strs := []string{"heooeh","hello", "a", "heoeh",""}

	for _,str := range strs{
		list := NewLinkedList()
		for _, c := range str{
			list.InsertTail(string(c))
		}
		list.Print()
		t.Log(IsPalindrome1(list))
	}
}


func TestIsPalindrome2(t *testing.T) {
	strs := []string{"heooeh","hello", "a", "heoeh",""}

	for _,str := range strs{
		list := NewLinkedList()
		for _, c := range str{
			list.InsertTail(string(c))
		}
		list.Print()
		t.Log(IsPalindrome1(list))
	}
}