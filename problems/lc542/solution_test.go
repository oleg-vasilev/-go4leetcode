package lc541

import (
	"reflect"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Ex1",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{0, 0, 0},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		{
			name: "Ex2",
			args: args{
				mat: [][]int{
					{0, 0, 0},
					{0, 1, 0},
					{1, 1, 1},
				},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
		{
			name: "Ex3",
			args: args{
				mat: [][]int{
					{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
					{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
					{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
					{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
					{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
					{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
					{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
					{1, 0, 0, 0, 1, 1, 1, 1, 0, 1},
					{1, 1, 1, 1, 1, 1, 1, 0, 1, 0},
					{1, 1, 1, 1, 0, 1, 0, 0, 1, 1},
				},
			},
			want: [][]int{
				{1, 0, 1, 1, 0, 0, 1, 0, 0, 1},
				{0, 1, 1, 0, 1, 0, 1, 0, 1, 1},
				{0, 0, 1, 0, 1, 0, 0, 1, 0, 0},
				{1, 0, 1, 0, 1, 1, 1, 1, 1, 1},
				{0, 1, 0, 1, 1, 0, 0, 0, 0, 1},
				{0, 0, 1, 0, 1, 1, 1, 0, 1, 0},
				{0, 1, 0, 1, 0, 1, 0, 0, 1, 1},
				{1, 0, 0, 0, 1, 2, 1, 1, 0, 1},
				{2, 1, 1, 1, 1, 2, 1, 0, 1, 0},
				{3, 2, 2, 1, 0, 1, 0, 0, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix()")
				t.Errorf("got :")
				for _, a := range got {
					t.Errorf("%v\n", a)
				}
				t.Errorf("want :")
				for _, a := range tt.want {
					t.Errorf("%v\n", a)
				}
			}
		})
	}
}
