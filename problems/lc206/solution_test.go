package lc206

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Ex#1",
			args: args{
				head: newLinkedListFromSlice([]int{1, 2, 3, 4, 5}),
			},
			want: newLinkedListFromSlice([]int{5, 4, 3, 2, 1}),
		},
		{
			name: "Ex#2",
			args: args{
				head: newLinkedListFromSlice([]int{1, 2}),
			},
			want: newLinkedListFromSlice([]int{2, 1}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func newLinkedListFromSlice(slice []int) *ListNode {
	if slice == nil || len(slice) == 0 {
		return nil
	}
	head := &ListNode{slice[0], nil}
	p := head
	for i := 1; i < len(slice); i++ {
		p.Next = &ListNode{slice[i], nil}
		p = p.Next
	}
	return head
}
