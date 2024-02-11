package lc278

import "testing"

func TestFirstBadVersion(t *testing.T) {
	cases := []struct {
		desc     string
		versions int
		blackbox versionBlackBox
		expected int
	}{
		{"Ex1", 10, versionBlackBox{badVersion: 5}, 5},
		{"Ex2", 1, versionBlackBox{badVersion: 1}, 1},
	}
	for _, tc := range cases {
		actual := firstBadVersion(tc.versions, tc.blackbox)
		if actual != tc.expected {
			t.Fatalf("%s: expected: %d got: %d",
				tc.desc, tc.expected, actual)
		}
	}
}
