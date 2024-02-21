package lc274

import (
	"testing"
)

func Test_hIndex(t *testing.T) {
	type args struct {
		citations []int
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
				[]int{3, 0, 6, 1, 5},
			},
			want: 3,
		},
		{
			name: "Ex2",
			args: args{
				[]int{1, 3, 1},
			},
			want: 1,
		},
		{
			name: "Ex3",
			args: args{
				[]int{100},
			},
			want: 1,
		},
		{
			name: "Ex4",
			args: args{
				[]int{0},
			},
			want: 0,
		},
		{
			name: "Ex4",
			args: args{
				[]int{0, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hIndex(tt.args.citations); got != tt.want {
				t.Errorf("hIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
