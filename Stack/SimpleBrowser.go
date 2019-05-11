package Stack

import "fmt"

type Browser struct {
	forwardStack Stack
	backStack Stack
}

func NewBrowser() *Browser  {
	return &Browser{
		forwardStack: NewArrayStack(),
		backStack: NewLinkerListStack(),
	}
}

func (brow *Browser) CanForwad()bool{
	if brow.forwardStack.IsEmpty(){
		return false
	}
	return true
}

func (brow *Browser) CanBack() bool  {
	if brow.backStack.IsEmpty(){
		return false
	}
	return true
}

func (brow *Browser)Open(addr string){
	// 打开新网页后无法再前进
	fmt.Println("Open new address ", addr)
	brow.forwardStack.Flush()
}

func (brow *Browser)PushBack(addr string)  {
	brow.backStack.Push(addr)
}

func  (brow *Browser)Forward()  {
	if brow.forwardStack.IsEmpty(){
		return
	}
	top := brow.forwardStack.Pop()
	brow.backStack.Push(top)
	fmt.Println("Forwad to ", top)
}

func (brow *Browser)Back()  {
	if brow.backStack.IsEmpty(){
		return
	}
	top := brow.backStack.Pop()
	brow.forwardStack.Push(top)
	fmt.Println("Back to ", top)
}
