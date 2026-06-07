package main

import "math"

type result struct {
	min  int
	max  int
	mode int
}

func apiTimings(data []int) result {
	min, max, mode := math.MaxInt, 0, 0
	freq := make(map[int]int)

	for _, n := range data {
		freq[n] += freq[n] + 1

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}

		if freq[n] > freq[mode] {
			mode = n
		}
	}

	return result{
		min:  min,
		max:  max,
		mode: mode,
	}

}
