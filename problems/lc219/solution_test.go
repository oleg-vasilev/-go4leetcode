package lc219

import (
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
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
				[]int{1, 2, 3, 1},
				3,
			},
			want: true,
		},
		{
			name: "Ex2",
			args: args{
				[]int{1, 0, 1, 1},
				1,
			},
			want: true,
		},
		{
			name: "Ex3",
			args: args{
				nums: []int{1, 2, 3, 1, 2, 3},
				k:    2,
			},
			want: false,
		},
		{
			name: "Ex4",
			args: args{
				nums: []int{1, 2, 3, 1, 1, 2, 3},
				k:    0,
			},
			want: false,
		},
		{
			name: "Ex5",
			args: args{
				nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9},
				k:    3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNearbyDuplicate(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("containsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
