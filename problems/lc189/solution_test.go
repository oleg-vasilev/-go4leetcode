package lc189

import (
	"reflect"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Ex1", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}, []int{5, 6, 7, 1, 2, 3, 4}},
		{"Ex2", args{[]int{-1, -100, 3, 99}, 2}, []int{3, 99, -1, -100}},
		{"Ex3", args{[]int{-17}, 2}, []int{-17}},
		{"Ex4", args{[]int{1, 2, 3, 4, 5, 6, 7}, 3}, []int{5, 6, 7, 1, 2, 3, 4}},
		{"Ex5", args{[]int{1, 2, 3, 4, 5, 6, 7}, 5}, []int{3, 4, 5, 6, 7, 1, 2}},
		{"Ex6", args{[]int{1, 2, 3}, 2}, []int{2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.args.nums, tt.args.k)
			if !reflect.DeepEqual(tt.args.nums, tt.want) {
				t.Errorf("rotate() = %v, want %v", tt.args.nums, tt.want)
			}
		})
	}
}
