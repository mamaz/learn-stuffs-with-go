package tokenbucket

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mamaz/token_bucket/cache"
)

// REFILL unit minutes
const RU_MINUTES = "minutes"
const RU_SECONDS = "seconds"

type Bucket struct {
	TokenNumber int
	refillUnit  string
	cache       cache.CacheI
	bucketId    string
}

func NewBucket(tokenNumber int, refillUnit string, cache cache.CacheI) *Bucket {
	return &Bucket{
		TokenNumber: tokenNumber,
		refillUnit:  refillUnit,
		cache:       cache,
		bucketId:    uuid.NewString(),
	}
}

func (b *Bucket) Start() {
	channel := make(chan int)

	go func() {
		for {
			switch b.refillUnit {
			case RU_MINUTES:
				time.Sleep(time.Duration(b.TokenNumber) * time.Minute)
			case RU_SECONDS:
				time.Sleep(time.Duration(b.TokenNumber) * time.Second)
			}
			fmt.Println("executing..")
			channel <- 1
		}
	}()

	for c := range channel {
		fmt.Println("check if we need refilling")

		if b.isEmpty() {
			fmt.Println("empty.. refilling")
			b.refill(c)
		}
	}
}

func (b *Bucket) refill(flag int) {
	b.cache.Set(b.bucketId, b.refillUnit)
}

func (b *Bucket) decrementToken() {
	counter := b.cache.Get(b.bucketId).(int)

	b.cache.Set(b.bucketId, counter-1)
}

func (b *Bucket) isEmpty() bool {
	return b.cache.Get(b.bucketId) == nil
}

func (b *Bucket) HandleRequest(payload interface{}) (bool, error) {
	if b.isEmpty() {
		rate := fmt.Sprintf("%v requests / %v", b.TokenNumber, b.refillUnit)
		return true, fmt.Errorf("request is above %v", rate)
	}

	b.decrementToken()

	return false, nil
}
