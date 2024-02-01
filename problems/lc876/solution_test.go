package lc876

import (
	"reflect"
	"testing"
)

func Test_middleNode(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Ex1",
			args: args{head: newLinkedListFromSlice([]int{1, 2, 3, 4, 5})},
			want: newLinkedListFromSlice([]int{3, 4, 5}),
		},
		{
			name: "Ex2",
			args: args{head: newLinkedListFromSlice([]int{1, 2, 3, 4, 5, 6})},
			want: newLinkedListFromSlice([]int{4, 5, 6}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := middleNode(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("middleNode() = %v, want %v", got, tt.want)
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
