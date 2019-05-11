package Stack

type Stack interface {
	Push(v interface{})
	Pop() interface{}
	IsEmpty() bool
	Flush()
	Print()
}
