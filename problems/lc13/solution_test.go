package lc13

import (
	"testing"
)

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	type testCase struct {
		name string
		args args
		want int
	}
	tests := []testCase{
		{
			name: "Ex1",
			args: args{
				"III",
			},
			want: 3,
		},
		{
			name: "Ex2",
			args: args{
				"LVIII",
			},
			want: 58,
		},
		{
			name: "Ex3",
			args: args{
				"MCMXCIV",
			},
			want: 1994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
