package nongenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddInt(t *testing.T) {
	tests := []int{1, 1, 1, 1}

	result := AddInt(tests)

	assert.Equal(t, 4, result)
}
