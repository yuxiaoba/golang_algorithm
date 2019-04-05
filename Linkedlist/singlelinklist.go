package Linkedlist

import (
	"github.com/pkg/errors"
	"fmt"
)

type ListNode struct {
	next *ListNode
	value interface{}  // 空接口适应不同类型的值
}

type LinkedList struct {
	head *ListNode
	length int
}

// 创建新链表
func NewLinkedList() *LinkedList  {
	return &LinkedList{NewListNode(0), 0}
}

// 创建新节点
func NewListNode(v interface{}) *ListNode  {
	return &ListNode{nil, v}
}

// 获取下一个节点
func (list *ListNode) GetNext() *ListNode  {
	return list.next
}

// 获取节点的值
func (list *ListNode) GetValue() interface{}  {
	return list.value
}

// 在某个节点后面插入节点
func (list *LinkedList) InsertAfter(p *ListNode,
	v interface{}) bool{
	if nil == p {
		return false
	}
	newNode := NewListNode(v)
	oldNext := p.next   // 先保存p原来指向的节点以防丢失
	p.next = newNode
	newNode.next = oldNext
	list.length++
	return true
}

// 在某个节点前面插入节点
func (list *LinkedList) InsertBefore (p *ListNode,
	v interface{}) bool  {
	if p == nil{
		return false
	}
	cur := list.head.next
	pre := list.head

	for cur != nil{
		if cur == p{
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	newNode := NewListNode(v)
	pre.next = newNode
	newNode.next = p
	list.length ++
	return true
}

// 在链表头部插入节点
func (list *LinkedList) InsertHead( v interface{}) bool  {
	return list.InsertAfter(list.head, v)
}

// 在链表尾部插入节点
func (list *LinkedList) InsertTail ( v interface{}) bool  {
	cur := list.head
	for cur.next != nil{
		cur = cur.next
	}
	return list.InsertAfter(cur, v)
}

// 删除某个节点
func (list *LinkedList) DeleteNode(p *ListNode) bool  {
    if p == nil{
    	return false
	}
	cur := list.head.next
	pre := list.head

	for cur != nil{
		if cur == p{
			break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil {
		return false
	}
	pre.next = p.next
	p = nil
	list.length --
	return true
}

// 寻找节点
func (list *LinkedList) SearchNode(p *ListNode) ( *ListNode , error)  {
	if p == nil{
		return NewListNode(nil), errors.New("No Such Node")
	}
	cur := list.head.next
    pre := list.head

    for cur != nil{
    	if cur == p{
    		break
		}
		pre = cur
		cur = cur.next
	}
	if cur == nil{
		return NewListNode(nil), errors.New("Not Such ListNode")
	}
	return pre, nil
}

// 打印链表
func (list *LinkedList) Print() {
	cur := list.head.next
	format := ""
	for nil != cur{
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if cur != nil{
			format += "->"
		}
	}
	fmt.Println(format)
}