package backtracking

func wordSearch(board [][]byte, word string) bool {

	index := 0

	var find func(i, j, index int) bool

	find = func(i, j, index int) bool {
		if index == len(word) { // word found
			return true
		}

		if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
			return false
		}

		if board[i][j] != word[index] { // not the letter we want
			return false
		}

		curr := board[i][j]
		board[i][j] = '#' // mark visited

		if find(i, j+1, index+1) {
			return true
		}
		if find(i, j-1, index+1) {
			return true
		}
		if find(i+1, j, index+1) {
			return true
		}
		if find(i-1, j, index+1) {
			return true
		}

		board[i][j] = curr // restore visited

		return false
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if find(i, j, index) {
				return true
			}
		}
	}

	return false
}
