package main

import (
	"fmt"
	"sync"
)

func produce(label string, nums []int, out chan int) {
	for _, n := range nums {
		fmt.Printf("%s producing: %d\n", label, n)
		out <- n
	}
	close(out)
}

func fanIn(in1, in2 <-chan int, out chan int) {
	// Use wait group and goroutines to merge
	var wg sync.WaitGroup

	wg.Add(1)
	go consume(in1, out, &wg)

	wg.Add(1)
	go consume(in2, out, &wg)

	go func() {
		wg.Wait()
		close(out)
	}()
}

func consume(in <-chan int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range in {
		out <- val
	}
}

func main() {
	odd := make(chan int)
	even := make(chan int)
	merged := make(chan int)

	go produce("odd", []int{1, 3, 5, 7}, odd)
	go produce("even", []int{2, 4, 6, 8}, even)

	go fanIn(odd, even, merged)

	// Consume merged channel
	for val := range merged {
		fmt.Println("Received:", val)
	}

	fmt.Println("All done!")
}
