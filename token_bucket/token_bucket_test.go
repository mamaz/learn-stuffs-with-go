package tokenbucket

import (
	"testing"

	"github.com/mamaz/token_bucket/cache"
)

func TestShouldBeAbleTo(t *testing.T) {
	inMemory := cache.New()

	b := NewBucket(2, RU_SECONDS, inMemory)
	b.Start()
}
