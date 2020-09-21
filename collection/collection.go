package collection

type Collection interface {
	IsEmpty() bool
	Size() int
	Clear()
	Contains(e interface{}) bool
}
