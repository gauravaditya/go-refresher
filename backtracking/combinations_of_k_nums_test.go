package backtracking

import (
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	tests := []struct {
		n, k     int
		expected [][]int
	}{
		{
			n: 4, k: 2,
			expected: [][]int{
				{1, 2}, {1, 3}, {1, 4},
				{2, 3}, {2, 4},
				{3, 4},
			},
		},
		// {
		// 	n: 1, k: 1,
		// 	expected: [][]int{{1}},
		// },
		// {
		// 	n: 5, k: 3,
		// 	expected: [][]int{
		// 		{1, 2, 3}, {1, 2, 4}, {1, 2, 5},
		// 		{1, 3, 4}, {1, 3, 5}, {1, 4, 5},
		// 		{2, 3, 4}, {2, 4, 5}, {2, 3, 5},
		// 		{3, 4, 5},
		// 	},
		// },
	}
	for _, tt := range tests {
		got := combine(tt.n, tt.k)
		sort2DSlice(got)
		sort2DSlice(tt.expected)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("combine(%d, %d) = %v, want %v", tt.n, tt.k, got, tt.expected)
		}
	}
}
