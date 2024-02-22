package lc11

import (
	"testing"
)

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
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
				[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
		{
			name: "Ex2",
			args: args{
				[]int{1, 1},
			},
			want: 1,
		},
		{
			name: "Ex3",
			args: args{
				[]int{1, 3, 2, 5, 25, 24, 5},
			},
			want: 24,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
