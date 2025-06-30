package backtracking

import (
	"fmt"
	"strings"
)

func wordSearch(board [][]byte, word string) bool {
	var search func(path []byte, i, j int)

	var result bool

	// delta := 1

	search = func(path []byte, x, y int) {
		if string(path) == word || result {
			result = true
			return
		}

		if x < 0 || y < 0 {
			return
		}

		if !strings.HasPrefix(word, string(path)) {
			return
		}

		delta := 1
		visited := make(map[[2]int]bool)

		fmt.Println("entering:", string(path))
		for i := x; i < len(board); i++ {
			for j := y; j < len(board[x]) && j >= 0; {
				visited[[2]int{i, j}] = true
				curr := board[i][j]

				fmt.Printf("adding: [%d,%d], val: %s\n", i, j, string(curr))
				path = append(path, board[i][j])

				j += delta
				search(path, i, j)

				if !strings.HasPrefix(word, string(path)) {
					fmt.Printf("removing: [%d,%d], val: %s\n", i, j, string(curr))
					path = path[:len(path)-1]
				}

				if visited[[2]int{i, j}] {
					continue
				} else {
					delta *= -1
				}
			}
		}
	}

	search([]byte{}, 0, 0)
	return result
}
