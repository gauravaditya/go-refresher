package backtracking

import (
	"fmt"
	"strings"
)

func generateBinary(n int) []string {
	var backtrack func(path string)

	result := []string{}
	backtrack = func(path string) {
		indent := strings.Repeat("  ", len(path))
		fmt.Println(indent+"Entering:", path)

		if len(path) == n {
			fmt.Println(indent+"Exiting:", path)
			result = append(result, path)
			return
		}

		backtrack(path + "0")
		backtrack(path + "1")
	}

	backtrack("")

	return result
}
