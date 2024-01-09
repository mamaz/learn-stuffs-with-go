package main

import "fmt"

func main() {
	bufchan := make(chan string, 2)

	bufchan <- "one"
	bufchan <- "two"
	fmt.Println("get bufchan: ", <-bufchan)
	bufchan <- "three"

	for i := 0; i < 2; i++ {
		fmt.Println(<-bufchan)
	}
}
