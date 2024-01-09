package functions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScanNumOfWords(t *testing.T) {
	numOfWords, err := CountNumOfWords("../test.txt")
	assert.Nil(t, err)
	assert.Equal(t, 58164, numOfWords)
}
