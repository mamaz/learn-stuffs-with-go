// Package heap contains Heap implementation
package heap

type Heap interface {
	Push(value int) int
	PushMany(values []int)
	Pop() int
	GetRootValue() int
	Count() int
	Print()
}
