package Stack

import "testing"

func TestBrowser(t *testing.T) {
	b := NewBrowser()
	b.PushBack("www.qq.com")
	b.PushBack("www.baidu.com")
	b.PushBack("www.hao123.com")
	b.backStack.Print()
	b.forwardStack.Print()

	if b.CanBack(){
		b.Back()
	}
	if b.CanBack(){
		b.Back()
	}
	b.backStack.Print()
	b.forwardStack.Print()

	if b.CanForwad(){
		b.Forward()
	}
	b.backStack.Print()
	b.forwardStack.Print()

	b.Open("www.yuxiaoba.com")
	if b.CanForwad(){
		b.Forward()
	}
	b.backStack.Print()
	b.forwardStack.Print()


}