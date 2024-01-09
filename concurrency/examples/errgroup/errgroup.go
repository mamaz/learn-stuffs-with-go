package main

import (
	"context"
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"

	u "concurrency/util"
)

func main() {
	g, _ := errgroup.WithContext(context.Background())
	iterations := 1000

	// result := make([]string, 0)
	resultMap := map[int]string{}

	// concurrently get URLs
	// Note:
	// this might cause race conditions if we have lots of iteration and we're using mutable data structure
	// if you have values to return to, use channels to pass the values, don't mutate
	for i := 0; i < iterations; i++ {
		idx := i
		g.Go(func() error {
			url, err := u.GetRandomURLSometimesReturnError()
			if err != nil {
				return err
			}

			// result = append(result, url)
			resultMap[idx] = url

			return nil
		})
	}

	for i := 0; i < iterations; i++ {
		idx := i
		g.Go(func() error {
			url, err := u.GetRandomURLSometimesReturnError()
			if err != nil {
				return err
			}

			// result = append(result, url)
			resultMap[idx] = url

			return nil
		})
	}

	// return error and stops everything when one of g.Go returns one error
	if err := g.Wait(); err != nil {
		log.Printf("error on execution: %+v", err)
	}

	// fmt.Println("result: ", result)
	fmt.Println("result: ", resultMap)
	// fmt.Println("total: ", len(result))
	fmt.Println("total: ", len(resultMap))
}
