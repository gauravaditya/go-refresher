package backtracking

import (
	"fmt"
	"maps"
	"practise-go/datatype/trie"
	"slices"
)

func findWords(board [][]byte, words []string) []string {
	root := trie.New()

	for _, word := range words {
		root.Insert(word)
	}

	found := make(map[string]string)

	var find func(i, j int, node *trie.Node, current []byte)

	find = func(i, j int, node *trie.Node, current []byte) {
		if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
			return
		}

		fmt.Println("current item: ->", string(board[i][j]))
		if !node.StartsWith(string(current)) {
			return
		}

		if node.Search(string(current)) {
			found[string(current)] = string(current)
		}

		temp := board[i][j]
		current = append(current, board[i][j])

		board[i][j] = '#' // mark visited

		fmt.Println("current: ->", string(current))
		find(i, j+1, node, current)
		find(i, j-1, node, current)
		find(i+1, j, node, current)
		find(i-1, j, node, current)

		board[i][j] = temp
		current = current[:len(current)-1]
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			find(i, j, root, []byte{})
		}
	}

	return slices.Collect(maps.Values(found))
}
