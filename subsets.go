package main

func subsets(nums []int) [][]int {
	subsets := make([][]int, 0)
	subsets = append(subsets, []int{}) // Start with the empty subset

	for _, v := range nums {
		// newSubset := []int{v}

		for _, subset := range subsets {
			subsets = append(subsets, append(subset, v)) // Add the current number to each existing subset
		}

		// subsets = append(subsets, newSubset)
	}

	return subsets // Start with the empty subset
}

func subsetsRecursive(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}} // Base case: return the empty subset
	}

	first := nums[0]
	rest := nums[1:]

	subsetsWithoutFirst := subsetsRecursive(rest)
	subsetsWithFirst := make([][]int, 0)

	for _, subset := range subsetsWithoutFirst {
		newSubset := append([]int{first}, subset...) // Add the first element to each subset
		subsetsWithFirst = append(subsetsWithFirst, newSubset)
	}

	return append(subsetsWithoutFirst, subsetsWithFirst...) // Combine both subsets
}
