package main

import "testing"

func TestLengthOfLongestSubstringAlt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abcdef", 6},
		{"aab", 2},
		{"jbpnbwwd", 4},
		{"jbpnbwewd", 5},
		{"abba", 2},
	}

	for _, test := range tests {
		result := lengthOfLongestSubstringAlt(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}
