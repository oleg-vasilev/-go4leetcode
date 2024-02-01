package lc3

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
		{name: "Example#1", args: args{s: "abcabcbb"}, want: 3}, // abc
		{name: "Example#2", args: args{s: "bbbbb"}, want: 1},    // b
		{name: "Example#3", args: args{s: "pwwkew"}, want: 3},   // wke
		{name: "Example#4", args: args{s: " "}, want: 1},        // " "
		{name: "Example#5", args: args{s: "a"}, want: 1},        // "a"
		{name: "Example#6", args: args{s: "dvdf"}, want: 3},     // "vdf"
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
