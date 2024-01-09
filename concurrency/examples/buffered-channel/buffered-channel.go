package main

import "fmt"

func main() {
	bufchan := make(chan string, 2)

	// insert data to goroutine one by one
	// until it hits the buffer limit
	bufchan <- "one"
	bufchan <- "two"
	fmt.Println("get bufchan: ", <-bufchan) // need to
	bufchan <- "three"

	for i := 0; i < 2; i++ {
		fmt.Println(<-bufchan)
	}
}
