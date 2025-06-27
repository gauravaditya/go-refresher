package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

func TestSubsets(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		// {
		// 	nums:     []int{},
		// 	expected: [][]int{{}}, // empty set
		// },
		// {
		// 	nums:     []int{1},
		// 	expected: [][]int{{}, {1}},
		// },
		// {
		// 	nums:     []int{1, 2},
		// 	expected: [][]int{{}, {1}, {2}, {1, 2}},
		// },
		{
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}

	for _, tt := range tests {
		got := subsets(tt.nums)
		sort2DSlice(got)
		sort2DSlice(tt.expected)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("subsets(%v) = %v, want %v", tt.nums, got, tt.expected)
		}
	}
}

func sort2DSlice(s [][]int) {
	for _, inner := range s {
		sort.Ints(inner)
	}
	sort.Slice(s, func(i, j int) bool {
		if len(s[i]) != len(s[j]) {
			return len(s[i]) < len(s[j])
		}
		for k := range s[i] {
			if s[i][k] != s[j][k] {
				return s[i][k] < s[j][k]
			}
		}
		return false
	})
}
