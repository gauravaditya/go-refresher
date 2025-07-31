package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	inbound := make(chan int)
	outbound := make(chan int)
	var wg sync.WaitGroup

	// Create a single worker
	w1 := newWorker("worker1", inbound, outbound)

	// Send values into inbound channel
	go func() {
		for i := 0; i < 10; i++ {
			inbound <- i
		}
		close(inbound)
	}()

	// Process and print using worker
	for i := 0; i < 10; i++ {
		wg.Add(2)
		go w1.processInbound(&wg)
		go w1.printOutbound(&wg)
	}

	wg.Wait()
}

type worker struct {
	label string
	src   <-chan int
	dest  chan int
}

func (w *worker) processInbound(wg *sync.WaitGroup) {
	defer wg.Done()
	num, ok := <-w.src
	if !ok {
		return
	}
	time.Sleep(1 * time.Second)
	w.dest <- num
}

func (w *worker) printOutbound(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(w.label, ":", <-w.dest)
}

func newWorker(label string, src <-chan int, dest chan int) *worker {
	return &worker{
		label: label,
		src:   src,
		dest:  dest,
	}
}
