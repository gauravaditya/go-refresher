package backtracking

import (
	"fmt"
	"sort"
	"strings"
)

func subsetsWithDupes(nums []int) [][]int {
	subsets := make([][]int, 0)
	sort.Ints(nums)

	var backtrack func(path []int, start int)

	backtrack = func(path []int, start int) {
		indent := strings.Repeat("  ", len(path))
		seen := make(map[int]bool)

		subsets = append(subsets, append([]int{}, path...))

		for i := start; i < len(nums); i++ {
			if seen[nums[i]] {
				continue
			}
			seen[nums[i]] = true

			path = append(path, nums[i])

			fmt.Println(indent+"enter-->", path, "i=", start)
			backtrack(path, i+1)

			fmt.Println(indent+"exit<--", path, "i=", start)
			path = path[:len(path)-1]
		}

	}

	backtrack([]int{}, 0)

	return subsets
}
