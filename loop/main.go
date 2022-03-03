package main

import (
	"fmt"
	"time"
)

func main() {
	iteration := 100_000_000
	counter := 0

	start := time.Now().UTC()
	for i := 0; i < iteration; i++ {
		counter++
	}
	elapsed := time.Since(start)

	fmt.Printf("Counter: %v\n", counter)
	fmt.Printf("Elapsed: %v ms", elapsed.Milliseconds())
}
