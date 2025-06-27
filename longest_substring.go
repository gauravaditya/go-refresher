package main

func lengthOfLongestSubstring(s string) int {
	seen := make(map[rune]int)
	max, start := 0, 0
	for i, r := range s {
		if lastSeen, ok := seen[r]; ok && lastSeen >= start {
			start = lastSeen + 1
		}
		seen[r] = i

		if i-start+1 > max {
			max = i - start + 1
		}
	}

	return max
}