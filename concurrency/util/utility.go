package util

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// GetRandomURLWithChannel simulates getting random URL from an endpoint
func GetRandomURLWithChannel(urlChan chan string) {
	urls := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://www.jamtangan.com",
		"https://voila.id",
	}

	source := rand.NewSource(time.Now().UnixNano())
	randTime := rand.New(source).Intn(1000)
	// give sleep to simulate latency
	time.Sleep(time.Duration(randTime) * time.Millisecond)

	randIndex := rand.New(source).Intn(len(urls))
	urlChan <- urls[randIndex]
}

func GetRandomURL() string {
	urls := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://www.jamtangan.com",
		"https://voila.id",
	}

	source := rand.NewSource(time.Now().UnixNano())
	randTime := rand.New(source).Intn(1000)

	// give sleep to simulate latency
	time.Sleep(time.Duration(randTime) * time.Millisecond)

	// give random a seed
	randIndex := rand.New(source).Intn(len(urls))

	return urls[randIndex]
}

func GetRandomURLSometimesReturnError() (string, error) {
	urls := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://www.jamtangan.com",
		"https://voila.id",
	}

	source := rand.NewSource(time.Now().UnixNano())
	randTime := rand.New(source).Intn(1000)

	// give sleep to simulate latency
	time.Sleep(time.Duration(randTime) * time.Millisecond)

	// give random a seed
	randIndex := rand.New(source).Intn(len(urls))

	// randomised error
	errOrNo := []error{errors.New("error happens"), nil, errors.New("error happens"), nil, nil, nil, nil}
	randErrIndex := rand.New(source).Intn(len(errOrNo))

	return urls[randIndex], errOrNo[randErrIndex]
}

func GetRandomURLSometimesReturnErrorWithChannel(okchan chan string, errchan chan error) {
	urls := []string{
		"https://google.com",
		"https://yahoo.com",
		"https://www.jamtangan.com",
		"https://voila.id",
	}

	source := rand.NewSource(time.Now().UnixNano())
	randTime := rand.New(source).Intn(1000)

	// give sleep to simulate latency
	time.Sleep(time.Duration(randTime) * time.Millisecond)

	// give random a seed
	randIndex := rand.New(source).Intn(len(urls))

	// randomised error
	errOrNo := []error{errors.New("error happens"), nil, errors.New("error happens"), nil, nil, nil, nil}
	randErrIndex := rand.New(source).Intn(len(errOrNo))

	if errOrNo[randErrIndex] != nil {
		errchan <- errOrNo[randErrIndex]
		return
	}

	okchan <- urls[randIndex]
}

func DeferredFunc(wg *sync.WaitGroup, function func() string) {
	defer wg.Done()
	res := function()

	fmt.Println("result: ", res)
}

func AddWgDone(wg *sync.WaitGroup, function func() string) func() string {
	return func() string {
		defer wg.Done()

		res := function()
		fmt.Println("result: ", res)

		return res
	}
}
