package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 900*time.Millisecond)
	defer cancel()

	reschan := make(chan string)

	go CallAPI(ctx, reschan)

	select {
	case rep := <-reschan:
		fmt.Println("result", rep)
	case <-ctx.Done():
		fmt.Println("deadline exceeded!", ctx.Err())
	}
}

func CallAPI(ctx context.Context, reply chan string) {
	time.Sleep(1 * time.Second)
	reply <- `{"data": "ok"}`
}
