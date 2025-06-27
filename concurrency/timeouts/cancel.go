package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: cancelled\n", id)
			return
		default:
			fmt.Printf("Worker %d: working...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func cancelWithContext() {
	fmt.Println("------------->>>>>>>>>>>")
	fmt.Println("------------->>>>>>>>>>>")
	fmt.Println("Running: Context with cancel...")
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, 1)
	go worker(ctx, 2)

	time.Sleep(2 * time.Second) // time for work

	fmt.Println("caller: invoking cancel...")
	cancel()

	time.Sleep(1 * time.Second)
	fmt.Println("caller: exiting...")
}
