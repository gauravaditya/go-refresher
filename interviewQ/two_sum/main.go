package main

func twoSum(input []int, target int) []int {
	pair := make(map[int]int)

	for i, n := range input {
		if j, ok := pair[n]; ok {
			return []int{i, j}
		}

		pair[target-n] = i
	}

	return []int{}
}
