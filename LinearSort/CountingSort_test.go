package LinearSort

import "testing"

func TestCountingSort(t *testing.T) {
	a := []int{1,6,3,5,8,6,4}
	CountingSort(a)
	t.Log(a)
}
