package tokenbucket

import (
	"testing"

	"github.com/mamaz/token_bucket/cache"
	"github.com/stretchr/testify/assert"
)

func TestShouldBeAbleToPreventRequests(t *testing.T) {
	inMemory := cache.New()

	b := NewBucket(2, RU_SECONDS, inMemory)
	_, decrementChan := b.Start()

	b.HandleRequest("is oke", decrementChan)
	b.HandleRequest("is oke", decrementChan)

	// request should be rejected, because we have 2 request per seconds limiter
	isOk, err := b.HandleRequest("is oke", decrementChan)

	assert.Equal(t, false, isOk)
	assert.NotNil(t, err)
}

func TestShouldBeAbleToHandleRequestWithinTime(t *testing.T) {
	inMemory := cache.New()

	b := NewBucket(2, RU_SECONDS, inMemory)
	_, decrementChan := b.Start()

	b.HandleRequest("is oke", decrementChan)
	isOk, err := b.HandleRequest("is oke", decrementChan)

	assert.Equal(t, true, isOk)
	assert.Nil(t, err)
}
