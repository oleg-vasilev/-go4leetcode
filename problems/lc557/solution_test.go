package lc557

import "testing"

func Test_reverseWords(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Example 1", args{"Let's take LeetCode contest"}, "s'teL ekat edoCteeL tsetnoc"},
		{"Example 2", args{"God Ding"}, "doG gniD"},
		{"Example 3", args{"I love u"}, "I evol u"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
