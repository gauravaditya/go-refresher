package problems

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
)

// Helper to capture stdout
func captureOutput(f func()) string {
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = stdout
	buf.ReadFrom(r)
	return buf.String()
}

func TestFanOutRoundRobin_Basic(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	output := captureOutput(func() {
		fanOutRoundRobin(nums)
	})

	for _, n := range nums {
		expected := fmt.Sprintf("consuming: %d", n)
		if !strings.Contains(output, expected) {
			t.Errorf("expected output to contain %q, got %q", expected, output)
		}
	}
}

func TestFanOutRoundRobin_Empty(t *testing.T) {
	nums := []int{}
	output := captureOutput(func() {
		fanOutRoundRobin(nums)
	})

	if strings.TrimSpace(output) != "" {
		t.Errorf("expected no output for empty input, got %q", output)
	}
}

func TestFanOutRoundRobin_SingleElement(t *testing.T) {
	nums := []int{42}
	output := captureOutput(func() {
		fanOutRoundRobin(nums)
	})

	expected := "consuming: 42"
	if !strings.Contains(output, expected) {
		t.Errorf("expected output to contain %q, got %q", expected, output)
	}
}
