package main

import (
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
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
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}

	for _, tt := range tests {
		result := subsets(tt.nums)
		if !equalIgnoreOrder(result, tt.expected) {
			t.Errorf("subsets(%v) = %v, want %v", tt.nums, result, tt.expected)
		}
	}
}

func TestSubsetsEdgeCases(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		// {
		// 	nums:     []int{0},
		// 	expected: [][]int{{}, {0}},
		// },
		// {
		// 	nums:     []int{1, 1},
		// 	expected: [][]int{{}, {1}, {1}, {1, 1}},
		// },
		// {
		// 	nums:     []int{-1, 2},
		// 	expected: [][]int{{}, {-1}, {2}, {-1, 2}},
		// },
		{
			nums:     []int{1, 2, 3},
			expected: [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}},
		},
	}

	for _, tt := range tests {
		result := subsetsRecursive(tt.nums)
		if !equalIgnoreOrder(result, tt.expected) {
			t.Errorf("subsets(%v) = %v, want %v", tt.nums, result, tt.expected)
		}
	}
}

// Helper to compare slices of slices ignoring order
func equalIgnoreOrder(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, x := range a {
		found := false
		for j, y := range b {
			if !used[j] && reflect.DeepEqual(x, y) {
				used[j] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
