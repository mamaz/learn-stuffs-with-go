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
	bucketId := uuid.NewString()
	cache.Set(bucketId, tokenNumber)

	return &Bucket{
		TokenNumber: tokenNumber,
		refillUnit:  refillUnit,
		cache:       cache,
		bucketId:    bucketId,
	}
}

// Start thread safe token bucket
// returns refillChannel and decrement channel
func (b *Bucket) Start() (chan bool, chan int) {
	refillChan := make(chan bool)
	decrementChan := make(chan int)

	// start handlers for refill and decrement token in a bucket
	go func() {
		for {
			select {
			case <-refillChan:
				b.refill()
			case <-decrementChan:
				b.decrementToken()
			}
		}
	}()

	// start goroutine for refilling token
	go func() {
		for {
			switch b.refillUnit {
			case RU_MINUTES:
				time.Sleep(time.Duration(b.TokenNumber) * time.Minute)
			case RU_SECONDS:
				time.Sleep(time.Duration(b.TokenNumber) * time.Second)
			}

			refillChan <- true
		}
	}()

	return refillChan, decrementChan
}

func (b *Bucket) refill() {
	b.cache.Set(b.bucketId, b.TokenNumber)
}

func (b *Bucket) decrementToken() {
	counter := b.cache.Get(b.bucketId).(int)

	b.cache.Set(b.bucketId, counter-1)
}

func (b *Bucket) isEmpty() bool {
	value := b.cache.Get(b.bucketId)
	return value == nil || value.(int) == 0
}

func (b *Bucket) HandleRequest(payload interface{}, decrementChan chan int) (bool, error) {
	if b.isEmpty() {
		rate := fmt.Sprintf("%v requests / %v", b.TokenNumber, b.refillUnit)
		return false, fmt.Errorf("request is above %v", rate)
	}

	decrementChan <- 1

	return true, nil
}
