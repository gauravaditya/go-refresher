package main

import (
	"fmt"
	"time"
)

// You have two channels: fast and slow.
// A value might come from either,
// but if nothing comes within 2 seconds, you should print "timeout" and exit.

func produceInts(label string, count int, ch chan int, timeout bool) {
	defer close(ch)

	for i := 1; i <= count; i++ {
		fmt.Println("producing", i, "into", label)
		ch <- i
		if timeout {
			time.Sleep(time.Second * time.Duration(i))
		}
	}
}

func zeroValue() {
	ch := make(chan int)
	go func() {
		ch <- 0
		close(ch)
	}()

	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("closed channel zero value:", val)
			break
		}
		fmt.Println("received:", val)
	}
}

func timeout() {
	fmt.Println("Running: Timeout...")
	size := 10
	fast := make(chan int, size)
	slow := make(chan int, size)

	go produceInts("slow", size, slow, true)
	go produceInts("fast", size, fast, false)

exit:
	for {
		select {
		case val, ok := <-fast:
			if !ok {
				fast = nil
			}
			fmt.Println("fast received:", val)
		case val, ok := <-slow:
			if !ok {
				slow = nil
			}
			fmt.Println("slow received:", val)
		case <-time.After(2 * time.Second): // timeout triggers exit if no channel receives data in 2 seconds
			fmt.Println("Timeout!")
			break exit
		}
	}

	zeroValue() // to differentiate between an actual numeric zero and a zero value for int on closed channel
	fmt.Println("Exiting main!")
}

func main() {
	timeout()
	cancelWithContext()
}
