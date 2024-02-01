package lc784

import (
	"reflect"
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want []string
}{
	{
		name: "Ex1",
		args: args{s: "a1b2"},
		want: []string{"a1b2", "a1B2", "A1b2", "A1B2"},
	},
}

func Test_letterCasePermutation(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCasePermutation(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCasePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_letterCasePermutationBacktrack(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := letterCasePermutationBacktrack(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCasePermutationBacktrack() = %v, want %v", got, tt.want)
			}
		})
	}
}
