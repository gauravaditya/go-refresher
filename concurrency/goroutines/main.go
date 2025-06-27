package main

import (
	"fmt"
	"sync"
)

func printNumbers(count int, wg *sync.WaitGroup) {
	for i := 1; i <= count; i++ {
		fmt.Println(i)
	}

	wg.Done()
}

func printLetters(from, to rune, wg *sync.WaitGroup) {
	for i := int(from); i <= int(to); i++ {
		fmt.Println(string(rune(i)))
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go printNumbers(5, &wg)

	wg.Add(1)
	go printLetters('a', 'e', &wg)

	wg.Wait()
}
