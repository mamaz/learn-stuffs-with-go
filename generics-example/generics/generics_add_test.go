package generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddIntGenerics(t *testing.T) {
	tests := []int{1, 1, 1, 1}

	result := Add(tests)

	assert.Equal(t, 4, result)
}

func TestAddFloatGenerics(t *testing.T) {
	tests := []float64{1.0, 1.1, 1.2, 1.3}

	result := Add(tests)

	assert.Equal(t, 4.6, result)
}

func TestAddStringGenerics(t *testing.T) {
	tests := []string{"ba", "na", "na"}

	result := Add(tests)

	assert.Equal(t, "banana", result)
}
