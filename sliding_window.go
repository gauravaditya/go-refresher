package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	out := sliding_window(os.Args[1])
	if strconv.Itoa(out) == os.Args[2] {
		fmt.Println("Output matches: ", out)
	} else {
		fmt.Println("Output doesn't match: ", out)
	}
}

func sliding_window(input string) int {
	seen := make(map[rune]int)
	start, max := 0, 0

	for i, r := range input {
		if lastSeen, ok := seen[r]; ok && lastSeen > start {
			start = lastSeen + 1
		}

		seen[r] = i
		if i-start+1 > max {
			max = i - start + 1
		}
	}

	return max
}
