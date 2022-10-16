package main

import (
	"fmt"
	"time"
)

func main() {
	// sleep in 12 seconds

	schan := sleep(5)
	<-schan

	// it should block for some seconds
	fmt.Println("finished")
}

func sleep(someseconds int64) chan struct{} {
	schan := make(chan struct{})
	go func() {
		counter := 0
		for {
			time.Sleep(1 * time.Second)
			counter += 1

			fmt.Println(counter)

			if counter == int(someseconds) {
				break
			}
		}
		schan <- struct{}{}
	}()

	return schan
}
