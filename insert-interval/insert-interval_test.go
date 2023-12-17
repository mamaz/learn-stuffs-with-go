package insertinterval

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertIntervalMergePartial(t *testing.T) {
	existing := []Interval{
		NewInterval(1, 3, false),
		NewInterval(4, 5, false),
		NewInterval(5, 8, false),
	}

	result := insertInterval(existing, NewInterval(2, 10, false))

	assert.Equal(t, NewInterval(1, 10, false), result)
}

func TestInsertIntervalMergeAll(t *testing.T) {
	existing := []Interval{
		NewInterval(1, 3, false),
		NewInterval(6, 9, false),
		NewInterval(10, 12, false),
	}

	result := insertInterval(existing, NewInterval(0, 23, false))

	assert.Equal(t, NewInterval(1, 10, false), result)
}

func TestInsertIntervalInsert(t *testing.T) {

}

func TestInsertMiddle(t *testing.T) {
	result := insertAt[int]([]int{1, 2, 3}, 10, 1)

	assert.Equal(t, []int{1, 10, 2, 3}, result)
}

func TestInsertFront(t *testing.T) {
	result := insertAt[int]([]int{1, 2, 3}, 10, 0)

	assert.Equal(t, []int{10, 1, 2, 3}, result)
}

func TestInsertEnd(t *testing.T) {
	result := insertAt[int]([]int{1, 2, 3}, 10, 2)

	assert.Equal(t, []int{1, 2, 3, 10}, result)
}

func TestInsertEndOverIndex(t *testing.T) {
	result := insertAt[int]([]int{1, 2, 3}, 10, 4)

	assert.Equal(t, []int{1, 2, 3, 10}, result)
}
