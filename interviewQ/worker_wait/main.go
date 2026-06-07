package main

import "fmt"

func main() {
	workerCount := 4
	done := make(chan int)

	for i := range workerCount {
		go worker(i, done)
	}

	for range workerCount {
		fmt.Printf("Exited worker: %d\n", <-done)
	}
}

func worker(id int, out chan int) {
	out <- id
}
