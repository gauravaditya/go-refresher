package backtracking

import (
	"testing"
)

func TestFindWords_EmptyBoardAndWords(t *testing.T) {
	board := [][]byte{}
	words := []string{}
	got := findWords(board, words)
	if len(got) != 0 {
		t.Errorf("Expected empty result, got %v", got)
	}
}

func TestFindWords_EmptyBoard(t *testing.T) {
	board := [][]byte{}
	words := []string{"a", "b"}
	got := findWords(board, words)
	if len(got) != 0 {
		t.Errorf("Expected empty result, got %v", got)
	}
}

func TestFindWords_EmptyWords(t *testing.T) {
	board := [][]byte{
		{'a', 'b'},
		{'c', 'd'},
	}
	words := []string{}
	got := findWords(board, words)
	if len(got) != 0 {
		t.Errorf("Expected empty result, got %v", got)
	}
}

func TestFindWords_NonEmptyInputs(t *testing.T) {
	board := [][]byte{
		{'o', 'a', 'a', 'n'},
		{'e', 't', 'a', 'e'},
		{'i', 'h', 'k', 'r'},
		{'i', 'f', 'l', 'p'},
	}
	words := []string{"oath", "pea", "eat", "rain"}
	got := findWords(board, words)
	if len(got) != 3 {
		t.Errorf("Expected 3 results, got %v", got)
	}
}
func TestFindWords_NoWordsFound(t *testing.T) {
	board := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
	}
	words := []string{"xyz", "uvw"}
	got := findWords(board, words)
	if len(got) != 0 {
		t.Errorf("Expected no results, got %v", got)
	}
}

func TestFindWords_SingleWordFound(t *testing.T) {
	board := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
	}
	words := []string{"abc"}
	got := findWords(board, words)
	if len(got) != 1 || got[0] != "abc" {
		t.Errorf("Expected 1 result with word 'abc', got %v", got)
	}
}

func TestFindWords_MultipleWordsFound(t *testing.T) {
	board := [][]byte{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
	}
	words := []string{"abc", "def", "gh"}
	got := findWords(board, words)
	if len(got) != 2 || (got[0] != "abc" && got[1] != "def") {
		t.Errorf("Expected 2 results with words 'abc' and 'def', got %v", got)
	}
}
