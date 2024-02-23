package lc134

import (
	"testing"
)

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
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
				[]int{1, 2, 3, 4, 5},
				[]int{3, 4, 5, 1, 2},
			},
			want: 3,
		},
		{
			name: "Ex2",
			args: args{
				[]int{2, 3, 4},
				[]int{3, 4, 3},
			},
			want: -1,
		},
		{
			name: "Ex3",
			args: args{
				[]int{3, 3, 4},
				[]int{3, 4, 4},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}
