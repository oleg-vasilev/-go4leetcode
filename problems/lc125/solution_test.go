package lc125

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	type testCase struct {
		name string
		args args
		want bool
	}
	tests := []testCase{
		{
			name: "Ex1",
			args: args{
				"A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "Ex2",
			args: args{
				"race a car",
			},
			want: false,
		},
		{
			name: "Ex3",
			args: args{
				" ",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
