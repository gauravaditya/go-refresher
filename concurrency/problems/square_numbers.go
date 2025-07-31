package problems

import (
	"fmt"
	"sync"
)

func generate(nums []int) <-chan int {
	c := make(chan int, len(nums))

	go func() {
		for _, num := range nums {
			c <- num
		}
		close(c)
	}()

	return c
}

func main() {
	fanOutRoundRobin([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func fanOutRoundRobin(nums []int) {
	out := make(chan int)
	var wg sync.WaitGroup

	workerCount := 3
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go process(i, workerCount, nums, out, &wg)
	}

	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(out)
	}(&wg)

	consumer(out)
}

func process(id, workerCount int, nums []int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := id; i < len(nums); {
		out <- nums[i]
		i = i + workerCount
	}
}

func consumer(in <-chan int) {
	for n := range in {
		fmt.Println("consuming:", n)
	}
}
