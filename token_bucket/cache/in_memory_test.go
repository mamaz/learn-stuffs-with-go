package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetKeyWithValue(t *testing.T) {
	cache := New()
	cache.Set("test", "tist")

	value := cache.Get("test")

	assert.Equal(t, "tist", value.(string))

}

func TestGetKey(t *testing.T) {
	cache := New()
	cache.Set("test", "tist")

	value := cache.Get("test")

	assert.Equal(t, "tist", value.(string))

}

func TestGetKeyButNotFound(t *testing.T) {
	cache := New()
	cache.Set("test", "tist")

	value := cache.Get("t")

	assert.Equal(t, nil, value)
}

func TestDeleteData(t *testing.T) {
	cache := New()
	cache.Set("test", "tist")

	cache.Delete("test")

	assert.Equal(t, nil, cache.Get("test"))
}
