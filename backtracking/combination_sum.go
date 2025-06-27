package backtracking

// candidates := []int{2, 3, 6, 7}
// target := 7

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)

	var backtrack func(path []int, sum, starts int)

	backtrack = func(path []int, sum, starts int) {
		if sum == target {
			result = append(result, append([]int{}, path...))
			return
		}

		if sum > target {
			return
		}

		for i := starts; i < len(candidates); i++ {
			num := candidates[i]

			path = append(path, candidates[i])
			sum += num
			backtrack(path, sum, i)

			sum -= num
			path = path[:len(path)-1]
		}
	}

	backtrack([]int{}, 0, 0)
	return result
}
