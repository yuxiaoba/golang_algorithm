package Array

import (
	"github.com/pkg/errors"
	"fmt"
)

type Array struct {
	data []int   // 数组内数据
	len  int     // 数组长度
}

// 为数组初始化内存
func NewArray(capacity int) *Array{
	if capacity == 0{
		return nil
	}
	return &Array{
		data: make([]int, capacity,capacity),
		len: 0,
	}
}

// 获取数组长度
func (array *Array) Len() int  {
	return array.len
}

// 判断索引是否越界
func (array Array) isIndexOutofRange (index int) bool{
	if index >= cap(array.data){
		return true
	}
	return false
}

// 通过索引查找数组，索引范围[0, n-1]
func (array Array) Search (index int) (int, error)  {
	if array.isIndexOutofRange(index){
		return  0, errors.New("out of index range")
    }
    return array.data[index], nil
}

// 插入数值到索引index上
func (array *Array) Insert (index int, v int) error{
	if array.len == cap(array.data){
		return errors.New("full array")
	}
	if array.isIndexOutofRange(index){
		return  errors.New("out of index range")
	}
	for i := array.len; i> index; i--{
		array.data[i] = array.data[i-1]
	}
	array.data[index] = v
	array.len++
	return nil
}

// 从尾部插入值
func (array *Array)InsertToTail(v int) error{
	return array.Insert(array.len, v)
}

// 删除索引index上的值
func (array *Array)Delete(index int) (int , error) {
	if array.isIndexOutofRange(index){
		return  0, errors.New("out of index range")
	}
	v := array.data[index]
	for i := index; i < array.len-1; i++{
		array.data[i] = array.data[i+1]
	}
	array.len --

	return  v, nil
}

// 打印数组
func (array *Array)Print()  {
	var format string
	for i := 0; i < array.len; i++{
		format += fmt.Sprintf("|%+v", array.data[i])
	}
	fmt.Println(format)
}





