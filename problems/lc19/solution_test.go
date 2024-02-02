package lc19

import (
	"reflect"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head *ListNode
		n    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Ex1",
			args: args{head: newLinkedListFromSlice([]int{1, 2, 3, 4, 5}), n: 2},
			want: newLinkedListFromSlice([]int{1, 2, 3, 5}),
		},
		{
			name: "Ex2",
			args: args{head: newLinkedListFromSlice([]int{1}), n: 1},
			want: nil,
		},
		{
			name: "Ex3",
			args: args{head: newLinkedListFromSlice([]int{1, 2}), n: 1},
			want: newLinkedListFromSlice([]int{1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeNthFromEnd(tt.args.head, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newLinkedListFromSlice(slice []int) *ListNode {
	head := &ListNode{slice[0], nil}
	p := head
	for i := 1; i < len(slice); i++ {
		p.Next = &ListNode{slice[i], nil}
		p = p.Next
	}
	return head
}
