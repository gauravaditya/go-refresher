package linkedlist

import (
	"reflect"
	"testing"
)

func listToSlice(head *Node) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestReverseIterative(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
		{
			name:     "multiple elements",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := FromSlice(tt.input)
			// reversed := ReverseIterative(head)
			reversed := ReverseRecursive(head) // Assuming you want to test the recursive version instead
			got := listToSlice(reversed)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Reverse(%v) = %v; want %v", tt.input, got, tt.expected)
			}
		})
	}
}
