package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		input  []int
		target int
		want   []int
	}{
		{
			name:   "Test case 1",
			input:  []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Test case 2",
			input:  []int{2, 7, 11, 15},
			target: 22,
			want:   []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.input, tt.target)
			assert.Contains(t, got, tt.want[0])
			assert.Contains(t, got, tt.want[1])
		})
	}
}
