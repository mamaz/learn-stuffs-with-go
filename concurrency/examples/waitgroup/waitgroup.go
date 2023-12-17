package main

import (
	u "concurrency/util"
	"sync"
)

// This example shows doing concurrent calls using waitgroup
// cannot return value in this one
// to return value on a goroutine, you need to add channels to it
func main() {
	var wg sync.WaitGroup
	iteration := 10

	for i := 0; i < iteration; i++ {
		wg.Add(1)

		// add wg.Done in u.GetRandomURL
		fun := u.AddWgDone(&wg, u.GetRandomURL)

		// cannot return value, since goroutine can only return value through channel
		go fun()
	}

	// wait until all function is done
	wg.Wait()
}
