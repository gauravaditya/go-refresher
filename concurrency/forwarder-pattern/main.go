package main

import "fmt"

// The forwarder pattern is a concurrency design pattern
// that allows you to decouple the production of data from its consumption.
// It involves using a separate goroutine (the forwarder) to receive data
// from a producer and then forward it to one or more consumers.
// This can help improve performance and reduce contention between producers and consumers.

// In this example, we will implement a simple forwarder pattern where a producer
// generates data and sends it to a forwarder, which then forwards the data to multiple consumers.

func main() {
	merged := merge(
		worker(1),
		worker(10),
		worker(100),
	)

	for v := range merged {
		fmt.Println(v)
	}
}

func worker(id int) <-chan int {
	out := make(chan int)
	go func() {
		for i := range 3 {
			out <- (i + 1) * id
		}
		defer close(out)
	}()

	return out
}

func merge(channels ...<-chan int) <-chan int {
	merged := make(chan int)
	done := make(chan struct{}, len(channels))
	for _, ch := range channels {
		go forward(ch, merged, done)
	}

	go func() {
		for i := 0; i < len(channels); i++ {
			<-done
		}
		close(done)
		close(merged)
	}()
	return merged
}

func forward(ch <-chan int, merged chan int, done chan<- struct{}) {
	for v := range ch {
		merged <- v
	}
	done <- struct{}{}
}
