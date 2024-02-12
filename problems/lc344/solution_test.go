package lc344

import (
	"reflect"
	"testing"
)

func Test_reverseString(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{"Example#1", args{[]byte("hello")}, []byte("olleh")},
		{"Example#2", args{[]byte("Hannah")}, []byte("hannaH")},
		{"Example#2", args{[]byte("X")}, []byte("X")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.args.bytes)
			if !reflect.DeepEqual(tt.args.bytes, tt.want) {
				t.Errorf("reverseString() = %v, want %v", tt.args.bytes, tt.want)
			}
		})
	}
}
