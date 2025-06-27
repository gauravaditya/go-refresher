package main

import "testing"

func TestMinWindow(t *testing.T) {
	tests := []struct {
		s, t     string
		expected string
	}{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"aa", "aa", "aa"}, // edge case
		{"abc", "aaa", ""}, // edge case
		{"", "a", ""},
		{"a", "", ""},
		{"abc", "d", ""},
		{"abdecfab", "abc", "cfab"},
		{"aa", "aa", "aa"},
		{"bba", "ab", "ba"},
		{"cabwefgewcwaefgcf", "cae", "cwae"},
		{"xyz", "xyz", "xyz"},
		{"xyz", "zyx", "xyz"},
		{"aaflslflsldkalskaaa", "aaa", "aaa"},
	}

	for _, tt := range tests {
		result := minWindow(tt.s, tt.t)
		if result != tt.expected {
			t.Errorf("minWindow(%q, %q) = %q; want %q", tt.s, tt.t, result, tt.expected)
		}
	}
}
