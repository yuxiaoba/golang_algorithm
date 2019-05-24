package Sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	a :=[]int{1,5,6,7,3,1}
	MergeSort(a)
	t.Log(a)
}