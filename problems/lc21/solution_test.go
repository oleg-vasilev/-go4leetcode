package lc21

import (
	"reflect"
	"testing"
)

type args struct {
	list1 *ListNode
	list2 *ListNode
}

type testcase struct {
	name string
	args args
	want *ListNode
}

var testsTable = []testcase{
	{
		name: "Ex#1",
		args: args{
			list1: newLinkedListFromSlice([]int{1, 2, 4}),
			list2: newLinkedListFromSlice([]int{1, 3, 4}),
		},
		want: newLinkedListFromSlice([]int{1, 1, 2, 3, 4, 4}),
	},
	{
		name: "Ex#2",
		args: args{
			list1: newLinkedListFromSlice([]int{}),
			list2: newLinkedListFromSlice([]int{}),
		},
		want: newLinkedListFromSlice([]int{}),
	},
	{
		name: "Ex#3",
		args: args{
			list1: newLinkedListFromSlice([]int{}),
			list2: newLinkedListFromSlice([]int{0}),
		},
		want: newLinkedListFromSlice([]int{0}),
	},
}

func Test_mergeTwoLists(t *testing.T) {
	var tests []testcase
	copy(tests, testsTable)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoListsNoRecursion(t *testing.T) {
	var tests []testcase
	copy(tests, testsTable)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoListsNoRecursion(tt.args.list1, tt.args.list2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists2() = %v, want %v", got, tt.want)
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
