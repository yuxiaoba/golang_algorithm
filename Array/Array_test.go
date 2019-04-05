package Array

import "testing"

func TestArray_Insert(t *testing.T) {
	capacity := 10

	arr := NewArray(capacity)

	for i:=0; i<capacity-2; i++{
		err := arr.Insert(i, i+1)
		if nil != err{
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	arr.Insert(6, 100)
	arr.Print()

	arr.InsertToTail(666)
	arr.Print()

}

func TestArray_Delete(t *testing.T) {
	capacity := 10

	arr := NewArray(capacity)

	for i:=0; i<capacity; i++{
		err := arr.Insert(i, i+1)
		if nil != err{
			t.Fatal(err.Error())
		}
	}
	arr.Print()

    for i:=9; i>=0; i--{
    	_, err := arr.Delete(i)
    	if nil != err{
    		t.Fatal(err)
		}
		arr.Print()
	}

}

func TestArray_Search(t *testing.T) {
	capacity := 10

	arr := NewArray(capacity)

	for i:=0; i<capacity; i++{
		err := arr.Insert(i, i+1)
		if nil != err{
			t.Fatal(err.Error())
		}
	}
	arr.Print()

	t.Log(arr.Search(0))
	t.Log(arr.Search(2))
	t.Log(arr.Search(12))
}