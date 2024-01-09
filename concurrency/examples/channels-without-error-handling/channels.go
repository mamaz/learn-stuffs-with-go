package main

import (
	u "concurrency/util"
	"fmt"
)

// This example shows fething data asynchronously on by one using goroutine and channels
// without error handling
func main() {
	urlChan := make(chan string)
	iteration := 10

	// asynchronously calls a function
	for i := 0; i < iteration; i++ {
		go u.GetRandomURLWithChannel(urlChan)
	}

	result := []string{}

	// getting the results
	// channel is blocked, get the value and append the values to a slice
	for i := 0; i < iteration; i++ {
		result = append(result, <-urlChan)
	}

	fmt.Printf("result: %+v", result)
}
