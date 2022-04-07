package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoundIndex(t *testing.T) {

	index := IndexOf([]int{1, 2, 3}, 2)
	assert.Equal(t, 1, index)
}

func TestFoundIndexString(t *testing.T) {

	index := IndexOf([]string{"ba", "na", "na"}, "na")
	assert.Equal(t, 1, index)
}

func TestNotFoundIndex(t *testing.T) {

	index := IndexOf([]int{1, 2, 3}, 13)
	assert.Equal(t, -1, index)
}
