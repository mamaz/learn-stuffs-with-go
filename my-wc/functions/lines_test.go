package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNumOfLine(t *testing.T) {
	numOfLines, err := CountNumberOfLines("../test.txt")
	assert.Equal(t, nil, err)
	assert.Equal(t, 7145, numOfLines)
}
