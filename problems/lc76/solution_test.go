package lc76

import (
	"testing"
)

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	type testCase struct {
		name string
		args args
		want string
	}
	tests := []testCase{
		{
			name: "Ex1",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC", // The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
		},

		{
			name: "Ex2",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a", // The entire string s is the minimum window.
		},

		{
			name: "Ex3",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "", // Both 'a's from t must be included in the window. Since the largest window of s only has one 'a', return empty string.
		},
		{
			name: "Ex4",
			args: args{
				s: "ab",
				t: "a",
			},
			want: "a",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
