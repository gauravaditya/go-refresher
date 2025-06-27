package backtracking

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_generateBinary(t *testing.T) {
	tests := map[int][]string{
		// 1: {"0", "1"},
		2: {"00", "01", "10", "11"},
		// 3: {"000", "001", "010", "011", "100", "101", "110", "111"},
		// 4: {"0000", "0001", "0010", "0011", "0100", "0101", "0110", "0111", "1000", "1001", "1010", "1011", "1110", "1110", "1111", "1111"},
	}

	for n, expected := range tests {
		t.Run(fmt.Sprintf("n=%d", n), func(t *testing.T) {
			got := generateBinary(n)
			assert.Condition(t, func() bool {
				if len(got) != len(expected) {
					return false
				}
				for _, v := range expected {
					if !assert.Contains(t, got, v) {
						return false
					}
				}
				return true
			}, "Expected %v, got %v", expected, got)
		})
	}
}
