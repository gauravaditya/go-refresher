package backtracking

func subsets(nums []int) [][]int {
	subsets := make([][]int, 0)

	var backtrack func(path []int, start int)

	backtrack = func(path []int, start int) {
		subsets = append(subsets, append([]int{}, path...))

		if len(path) == len(nums) {
			return
		}

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i]) // Include the current number
			backtrack(path, i+1)
			path = path[:len(path)-1] // Backtrack: remove the last number
		}
	}

	backtrack([]int{}, 0)

	return subsets // Start with the empty subset
}
