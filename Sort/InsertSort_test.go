package Sort

import (
	"testing"
	"fmt"
)

func TestInsertSort(t *testing.T) {
	a := []int{1,3,5,4,6,3}
	fmt.Print("排序前")
	fmt.Println(a)
	InsertSort(a)
	fmt.Print("排序后")
	fmt.Print(a)
}
