package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountNumOfBytes(t *testing.T) {
	numOfBytes, err := CountNumOfBytes("../test.txt")
	assert.Nil(t, err)
	assert.Equal(t, 342190, numOfBytes)
}

func TestShouldReturnZeroBytesIfFileIsEmpty(t *testing.T) {
	numOfBytes, err := CountNumOfBytes("../empty.txt")
	assert.Nil(t, err)
	assert.Equal(t, 0, numOfBytes)
}

func TestShouldReturnErrorIfNoFileIsSpecified(t *testing.T) {
	_, err := CountNumOfBytes("")
	assert.NotNil(t, err)
}
