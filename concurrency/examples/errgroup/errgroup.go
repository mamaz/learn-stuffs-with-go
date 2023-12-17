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
	iterations := 10

	result := make([]string, 0)

	for i := 0; i < iterations; i++ {
		g.Go(func() error {
			url, err := u.GetRandomURLSometimesReturnError()
			if err != nil {
				return err
			}

			result = append(result, url)

			return nil
		})
	}

	// return error and stops everything when one of g.Go returns one error
	if err := g.Wait(); err != nil {
		log.Printf("error on execution: %+v", err)
	}

	fmt.Println("result: ", result)
	fmt.Println("total: ", len(result))
}
