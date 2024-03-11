package lc6

import (
	"testing"
)

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
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
				s:       "PAYPALISHIRING",
				numRows: 3,
			},
			want: "PAHNAPLSIIGYIR",
		},
		{
			name: "Ex2",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 4,
			},
			want: "PINALSIGYAHRPI",
		},
		{
			name: "Ex99",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 1,
			},
			want: "PAYPALISHIRING",
		},
		{
			name: "Ex100",
			args: args{
				s:       "PAYPALISHIRING",
				numRows: 2,
			},
			want: "PYAIHRNAPLSIIG",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
