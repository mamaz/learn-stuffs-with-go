package heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertMaxValue(t *testing.T) {
	initialValue := []int{
		1, 2, 3,
	}
	h := NewHeap(initialValue)
	h.Push(4)

	assert.Equal(t, 4, h.GetRootValue())
}

func TestInsertManyValue(t *testing.T) {
	initialValue := []int{
		1, 2, 3,
	}
	h := NewHeap(initialValue)

	assert.Equal(t, 3, h.GetRootValue())
	assert.Equal(t, []int{3, 2, 1}, h.values)
}

func TestInsertLessThanRootValue(t *testing.T) {
	initialValue := []int{
		12, 1, 3,
	}
	h := NewHeap(initialValue)
	h.Push(4)

	assert.Equal(t, 12, h.GetRootValue())
}

func TestInsertTheSameValue(t *testing.T) {
	initialValue := []int{
		12,
	}
	h := NewHeap(initialValue)
	h.Push(12)

	assert.Equal(t, 12, h.GetRootValue())
}

func TestPopRoot(t *testing.T) {
	initialValue := []int{
		12, 11, 10,
	}
	h := NewHeap(initialValue)

	result := h.Pop()

	assert.Equal(t, 12, result)
	assert.Equal(t, 2, h.Count())
}
