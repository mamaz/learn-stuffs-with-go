package main

import (
	"concurrency/util"
	"fmt"
)

func main() {
	iteration := 10

	urlchan := make(chan string)
	errchan := make(chan error)

	for i := 0; i < iteration; i++ {
		go util.GetRandomURLSometimesReturnErrorWithChannel(urlchan, errchan)
	}

	result := []string{}

	for i := 0; i < iteration; i++ {
		// print one by one as soon as it's available
		// if error happens just prints
		select {
		case url := <-urlchan:
			fmt.Printf("request is successful: %+v\n", url)
			result = append(result, url)
		case err := <-errchan:
			fmt.Printf("error occurs: %+v\n", err)
		}
	}

	fmt.Printf("result: %+v", result)
}
