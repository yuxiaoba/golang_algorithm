package Linkedlist

/*
法一，将前半段入栈，然后对比
*/

func IsPalindrome1(list *LinkedList) bool{
	len := list.length
	if len == 0{
		return false
	}
	if len == 1 {
		return true
	}
	s := make([]string, 0, len/2)
	cur := list.head
	for i:=1; i<=len; i++{
		cur = cur.next
        if len % 2 != 0 && i == (len/2+1){
        	continue
		}
		if i <= len /2{
			s = append(s, cur.GetValue().(string))
		}else {
			if s[len-i] != cur.GetValue().(string){
				return false
			}
		}

	}
	return true
}

/*
快慢指针定位中间节点
将前半部分节点逆序
前后半部分对比判断是否为回文
*/
func IsPalindrome2(list *LinkedList) bool  {
	len := list.length
	if len == 0{
		return false
	}
	if len == 1 {
		return true
	}

	var slow = list.head.next
	var fast = list.head.next

	pre := NewListNode(nil)
	tmp := NewListNode(nil)

	for fast != nil && fast.next != nil{
		fast = fast.next.next

		// 前半部分逆序
		tmp = slow.next  // tmp 作为中间节点，记录当前节点下一个节点的位置
		slow.next = pre  // 当前节点指向前一个节点
		pre = slow       // 指针后移
		slow = tmp       // 指针后移，处理下一个节点
	}

	//如果是单数，需要跳过中间的节点
	if fast !=nil{
		slow = slow.next
	}

	if slow != nil && pre != nil &&
		slow.GetValue() == pre.GetValue(){
			slow = slow.next  // 后半段
			pre = pre.next    // 前半段
	}
	return pre == nil && slow== nil
}