package Linkedlist

// 链表反转
func (list *LinkedList) Reverse() (*LinkedList) {
	if list.length < 2{
		return nil
	}
	var pre *ListNode = nil
	var tmp *ListNode = nil
	cur := list.head.next

	for cur != nil{
		tmp = cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	list.head.next = pre

	return list
}

// 判断单链表是否有环
func (list *LinkedList) IsHasCycle() bool {
	if list.head != nil{
		slow := list.head
		fast := list.head
		for fast != nil && fast.next != nil{
			slow = slow.next
			fast = fast.next.next
			if slow == fast{
				return true
			}
		}
	}
	return false
}

// 两个有序链表合并
func MergeSortedList(l1, l2 *LinkedList) *LinkedList  {
	if l1 == nil || l1.length < 1{
		return l2
	}
	if l2 == nil || l2.length < 1{
		return l1
	}

	list := NewLinkedList()
	cur := list.head
	curl1 := l1.head.next
	curl2 := l2.head.next

	for curl1 != nil && curl2 != nil{
		if curl1.GetValue().(int) > curl2.GetValue().(int){
			cur.next = curl2
			curl2= curl2.next
		}else {
			cur.next = curl1
			curl1 = curl1.next
		}
		cur = cur.next
	}

	if curl1 != nil{
		cur.next = curl1
	}else if curl2 != nil{
		cur.next = curl2
	}

	return list
}

// 删除倒数第N个节点
func (list *LinkedList) DeleteBottleN (n int ) bool  {
	if n <= 0 || list.length < 1 {
		return false
	}
	fast := list.head
	slow := list.head
	for i :=0; i < n && fast != nil ; i++{
		fast = fast.next  // 其中一个指针先走N步
	}

	if fast == nil{
		return false
	}

	for fast.next != nil{
		slow = slow.next
		fast = fast.next
	}
	slow.next = slow.next.next
	return true
}

