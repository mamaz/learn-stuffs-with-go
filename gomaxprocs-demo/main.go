package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)

	// setting GOMAXPROCS to 0 will do nothing and prints current settings
	fmt.Println("max parallelism", runtime.GOMAXPROCS(0))
}
