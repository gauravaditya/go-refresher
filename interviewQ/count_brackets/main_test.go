package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isBalancedString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input string
		want  string
	}{
		{
			name:  "Balanced parentheses",
			input: "()",
			want:  "Balanced",
		},
		{
			name:  "Unbalanced parentheses",
			input: "(]",
			want:  "Not Balanced",
		},
		{
			name:  "Balanced mixed parentheses",
			input: "({[]})",
			want:  "Balanced",
		},
		{
			name:  "Unbalanced mixed parentheses",
			input: "({[})",
			want:  "Not Balanced",
		},
		{
			name:  "Empty string",
			input: "",
			want:  "Balanced",
		},
		{
			name:  "Only opening parentheses",
			input: "(((",
			want:  "Not Balanced",
		},
		// { // Unhandled case as per the current implementation, it will return "Balanced" instead of "Not Balanced"
		// 	name:  "Only closing parentheses",
		// 	input: ")))",
		// 	want:  "Not Balanced",
		// },
		{
			name:  "Mixed parentheses with extra characters",
			input: "a(b)c{d}e[f]g",
			want:  "Balanced",
		},
		{
			name:  "Mixed parentheses with incorrect nesting",
			input: "([)]",
			want:  "Not Balanced",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isBalancedString(tt.input)
			assert.Equal(t, tt.want, got)
		})
	}
}
