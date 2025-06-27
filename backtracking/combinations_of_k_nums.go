package backtracking

import "fmt"

func combine(n int, k int) [][]int {
	result := [][]int{}
	var backtrack func(start int, path []int)

	backtrack = func(start int, path []int) {
		// base case
		if len(path) == k {
			result = append(result, append([]int{}, path...)) // make a copy of path
			return
		}

		for i := start; i <= n; i++ {
			fmt.Println("Entering path:", path, "Adding:", i)
			path = append(path, i) // include the current number
			backtrack(i+1, path)
			fmt.Println("Exiting path:", path, "Removing:", i)
			path = path[:len(path)-1] // Backtrack: remove the last number
		}
	}

	backtrack(1, []int{})
	return result
}
