package lc191

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Ex1",
			args: args{num: 11}, // 00000000000000000000000000001011
			want: 3,
		},
		{
			name: "Ex2",
			args: args{num: 128}, // 00000000000000000000000010000000
			want: 1,
		},
		{
			name: "Ex2",
			args: args{num: 4294967293}, // 11111111111111111111111111111101
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
