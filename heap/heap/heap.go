package heap

import (
	"fmt"
)

type MaxHeap struct {
	values []int
}

func NewHeap(initialValue []int) *MaxHeap {
	newHeap := &MaxHeap{
		values: []int{},
	}
	newHeap.PushMany(initialValue)

	return newHeap
}

func (h *MaxHeap) Push(value int) {
	// add new value at the end of the slice
	h.values = append(h.values, value)

	index := len(h.values) - 1
	parentIndex := index / 2

	for index > 0 && h.values[index] >= h.values[parentIndex] {
		// swap values if it's smaller than parent
		temp := h.values[parentIndex]
		h.values[parentIndex] = h.values[index]
		h.values[index] = temp

		index = parentIndex
		parentIndex = parentIndex / 2
	}
}

func (h *MaxHeap) PushMany(values []int) {
	for _, v := range values {
		h.Push(v)
	}
}

func (h *MaxHeap) GetRootValue() int {
	return h.values[0]
}

func (h *MaxHeap) Pop() int {
	result := h.values[0]

	h.values[0] = h.values[h.Count()-1]

	// remove last element
	h.values = h.values[:len(h.values)-1]

	index := 0
	leftChildIndex := 2*index + 1
	rightChildIndex := 2*index + 2

	for leftChildIndex < h.Count() || rightChildIndex < h.Count() {
		left := h.values[leftChildIndex]
		right := h.values[rightChildIndex]

		if left >= right {
			if left >= h.values[index] {
				h.values[index] = h.values[leftChildIndex]
				index = leftChildIndex
			}
		} else {
			if right >= h.values[index] {
				h.values[index] = h.values[rightChildIndex]
				index = rightChildIndex
			}
		}

		leftChildIndex = 2*index + 1
		rightChildIndex = 2*index + 2
	}

	return result
}

func (h *MaxHeap) Print() {
	fmt.Printf("%+v\n", h.values)
}

func (h *MaxHeap) Count() int {
	return len(h.values)
}
