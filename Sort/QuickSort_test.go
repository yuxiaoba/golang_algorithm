package Sort

import (
	"testing"
	"log"
)

func TestQucikSort(t *testing.T) {
	a := []int{7,6,8,3,1,4,5}
	QucikSort(a)
	log.Print(a)
}
