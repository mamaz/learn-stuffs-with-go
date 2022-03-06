package tokenbucket

import (
	"testing"

	"github.com/mamaz/token_bucket/cache"
	"github.com/stretchr/testify/assert"
)

func TestShouldBeAbleToPreventRequests(t *testing.T) {
	inMemory := cache.New()

	b := NewBucket(2, RU_SECONDS, inMemory)
	b.Start()

	b.HandleRequest("is oke")
	b.HandleRequest("is oke")

	// request should be rejected, because we have 2 request per seconds limiter
	isOk, err := b.HandleRequest("is oke")

	assert.Equal(t, false, isOk)
	assert.NotNil(t, err)
}

func TestShouldBeAbleToHandleRequestWithinTime(t *testing.T) {
	inMemory := cache.New()

	b := NewBucket(2, RU_SECONDS, inMemory)
	b.Start()

	b.HandleRequest("is oke")
	isOk, err := b.HandleRequest("is oke")

	assert.Equal(t, true, isOk)
	assert.Nil(t, err)
}
