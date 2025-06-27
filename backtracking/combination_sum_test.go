package backtracking

import (
	"reflect"
	"sort"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{
			candidates: []int{2, 3, 6, 7},
			target:     7,
			want:       [][]int{{2, 2, 3}, {7}},
		},
		{
			candidates: []int{1, 2, 3},
			target:     4,
			want:       [][]int{{1, 1, 1, 1}, {1, 1, 2}, {1, 3}, {2, 2}},
		},
		{
			candidates: []int{2, 5, 3, 6, 7},
			target:     8,
			want:       [][]int{{2, 2, 2, 2}, {2, 3, 3}, {2, 6}, {3, 5}},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := combinationSum(tt.candidates, tt.target)
			// Sort both slices to compare them
			for i := range got {
				sort.Ints(got[i])
			}
			for i := range tt.want {
				sort.Ints(tt.want[i])
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum(%v, %d) = %v, want %v", tt.candidates, tt.target, got, tt.want)
			}
		})
	}
}
