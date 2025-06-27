package main

import "testing"

func Test_lengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "t1",
			args: args{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "t2",
			args: args{
				s: "bbbbbbb",
			},
			want: 1,
		},
		{
			name: "t3",
			args: args{
				s: "pwwkew",
			},
			want: 3,
		},
		{
			name: "t4",
			args: args{
				s: "",
			},
			want: 0,
		},
		{
			name: "t4",
			args: args{
				s: "aab",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
		{"abcdef", 6},
		{"aab", 2},
		{"jbpnbwwd", 4},
		{"jbpnbwewd", 5},
		{"abba", 2},
	}

	for _, test := range tests {
		result := lengthOfLongestSubstring(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected %d but got %d", test.input, test.expected, result)
		}
	}
}
