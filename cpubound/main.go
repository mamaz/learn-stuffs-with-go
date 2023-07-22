package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum += 1
	}

	elapsed := time.Since(start)
	fmt.Println("Result: ", sum)
	fmt.Println("Elapsed: ", elapsed.Milliseconds(), " ms")
}
