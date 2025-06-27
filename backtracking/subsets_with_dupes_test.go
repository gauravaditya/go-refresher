package backtracking

import (
	"reflect"
	"testing"
)

func TestSubsetsWithDupes(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums:     []int{},
			expected: [][]int{{}},
		},
		{
			nums:     []int{1},
			expected: [][]int{{}, {1}},
		},
		{
			nums:     []int{1, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}},
		},
		{
			nums:     []int{1, 2, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
		{
			nums:     []int{2, 2, 2},
			expected: [][]int{{}, {2}, {2, 2}, {2, 2, 2}},
		},
		{
			nums:     []int{2, 1, 2},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {2, 2}, {1, 2, 2}},
		},
	}

	for _, tt := range tests {
		got := subsetsWithDupes(tt.nums)
		sort2DSlice(got)
		sort2DSlice(tt.expected)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("subsetsWithDupes(%v) = %v, want %v", tt.nums, got, tt.expected)
		}
	}
}
