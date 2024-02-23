package lc12

import (
	"testing"
)

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
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
				3,
			},
			want: "III",
		},
		{
			name: "Ex2",
			args: args{58},
			want: "LVIII",
		},
		{
			name: "Ex3",
			args: args{1994},
			want: "MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
