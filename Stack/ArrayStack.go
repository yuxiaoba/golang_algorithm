package Stack

import (
	"fmt"
)

// 用数组实现栈

type ArrayStack struct {
	data []interface{}
	top int
}

//新建栈
func NewArrayStack() *ArrayStack{
	return &ArrayStack{
		data: make([]interface{},0,32),  // 返回一个len为0，cap为32的切片
		top: -1,
	}
}

// 判断是否为空
func (array *ArrayStack) IsEmpty()bool{
	if array.top < 0{
		return true
	}
	return false
}

func (array *ArrayStack) Push ( v interface{})  {
    array.top += 1

	if array.top > len(array.data)-1{
		// 对未存放过数据的切片空间
    	array.data = append(array.data, v)
	}else {
		array.data[array.top] =v
	}

}

func (array *ArrayStack) Pop() (interface{})  {
	if array.top < 0{
		fmt.Println("Stack is empty")
		return nil
	}else {
		v := array.data[array.top]
		array.top -= 1
		return v
	}
}

func (array *ArrayStack)Print()  {
	if array.IsEmpty(){
		fmt.Println("Stack is empty")
	}else{
		for i:=array.top; i>=0; i--{
			fmt.Printf("%d ", array.data[i])
		}
	}
}