package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_apiTimings(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		data []int
		want result
	}{
		{
			name: "Test case 1",
			data: []int{1, 2, 3, 4, 5},
			want: result{min: 1, max: 5, mode: 1}, // Update with expected values.
		},
		{
			name: "Test case 2",
			data: []int{2, 2, 3, 3, 4},
			want: result{min: 2, max: 4, mode: 2}, // Update with expected values.
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := apiTimings(tt.data)
			assert.Equal(t, tt.want.min, got.min, "Min value mismatch")
			assert.Equal(t, tt.want.max, got.max, "Max value mismatch")
			assert.Equal(t, tt.want.mode, got.mode, "Mode value mismatch")
		})
	}
}
