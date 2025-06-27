package backtracking

func permutations(nums []int) [][]int {
	result := [][]int{}

	var backtrack func(path []int, used map[int]bool)

	backtrack = func(path []int, used map[int]bool) {
		if len(path) == len(nums) {
			copyPath := append([]int{}, path...)
			result = append(result, copyPath) // make a copy of path

			return
		}

		for _, num := range nums {
			if used[num] {
				continue
			}

			used[num] = true

			path = append(path, num)
			backtrack(path, used)

			used[num] = false
			path = path[:len(path)-1]
		}
	}

	backtrack([]int{}, map[int]bool{})

	return result
}
