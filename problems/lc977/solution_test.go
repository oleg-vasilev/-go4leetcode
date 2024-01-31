package lc977

import (
	"reflect"
	"testing"
)

func Test_sortedSquares(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Ex1", args{[]int{-4, -1, 0, 3, 10}}, []int{0, 1, 9, 16, 100}},
		{"Ex2", args{[]int{-7, -3, 2, 3, 11}}, []int{4, 9, 9, 49, 121}},
		{"Ex3", args{[]int{1, 3, 5}}, []int{1, 9, 25}},
		{"Ex4", args{[]int{-1, 3, 5, 6}}, []int{1, 9, 25, 36}},
		{"Ex5", args{[]int{-5, -1, 0, 3}}, []int{0, 1, 9, 25}},
		{"Ex6", args{[]int{-5, -3}}, []int{9, 25}},
		{"Ex7", args{[]int{-5, 0, 8}}, []int{0, 25, 64}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortedSquares(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
