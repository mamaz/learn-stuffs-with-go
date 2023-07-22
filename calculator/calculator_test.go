package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldBeAbleToAddSubtractWithoutBracket(t *testing.T) {
	result := Calculate("1 + 2 - 1")
	assert.Equal(t, 2, result)
}

func TestShouldBeAbleToAddWithoutBracket(t *testing.T) {
	result := Calculate("1 + 2 + 1")
	assert.Equal(t, 4, result)
}
