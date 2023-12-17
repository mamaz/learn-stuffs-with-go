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

	for i := 0; i < iteration; i++ {
		go u.GetRandomURLWithChannel(urlChan)
	}

	result := []string{}

	// print one by one as soon as it's available
	for i := 0; i < iteration; i++ {
		result = append(result, <-urlChan)
	}

	fmt.Printf("result: %+v", result)
}
