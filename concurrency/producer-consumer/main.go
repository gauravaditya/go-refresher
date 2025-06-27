package main

import (
	"fmt"
	"sync"
	"time"
)

func produceNumbers(from, to int, ch chan int, wg *sync.WaitGroup) {
	for i := from; i < to+1; i++ {
		fmt.Println("producing ->", i)
		ch <- i
	}

	close(ch)
	wg.Done()
}

func consumeNumbers(name string, ch chan int, wg *sync.WaitGroup) {
	for val := range ch {
		fmt.Println(name, "consuming ->", val)
		time.Sleep(1 * time.Second)
	}

	wg.Done()
}

func runConsumers(count int, ch chan int, wg *sync.WaitGroup) {
	for i := 0; i < count; i++ {
		name := fmt.Sprintf("consumer%d", i)
		wg.Add(1)
		go consumeNumbers(name, ch, wg)
	}
}

func main() {
	bufferedCh := make(chan int, 3)
	// unbufferedCh := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go produceNumbers(1, 10, bufferedCh, &wg)

	// runConsumers(1, unbufferedCh, &wg) // single consumer scenario with unbuffered chan
	// runConsumers(1, bufferedCh, &wg) // single consumer scenario with buffered chan
	runConsumers(2, bufferedCh, &wg) // multiple consumers, load balancing example
	wg.Wait()
}
