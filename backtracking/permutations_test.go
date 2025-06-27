package backtracking

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	tests := []struct {
		input    []int
		expected [][]int
	}{
		{
			input: []int{},
			expected: [][]int{
				{},
			},
		},
		{
			input: []int{1},
			expected: [][]int{
				{1},
			},
		},
		{
			input: []int{1, 2},
			expected: [][]int{
				{1, 2},
				{2, 1},
			},
		},
		{
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.input), func(t *testing.T) {
			got := permutations(tt.input)
			sort2DSlice(got)
			sort2DSlice(tt.expected)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}
