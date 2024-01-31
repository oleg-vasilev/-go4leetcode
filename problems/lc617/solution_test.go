package lc617

import (
	"strconv"
	"testing"
)

func Test_mergeTrees(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "Example#1",
			args: args{
				root1: binTreeFromSlice([]string{"1", "3", "2", "5"}, 0),
				root2: binTreeFromSlice([]string{"2", "1", "3", "nil", "4", "nil", "7"}, 0),
			},
			want: binTreeFromSlice([]string{"3", "4", "5", "5", "4", "nil", "7"}, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeTrees(tt.args.root1, tt.args.root2)
			if !TreesEqual(got, tt.want) {
				t.Errorf("mergeTrees() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func binTreeFromSlice(data []string, i int) *TreeNode {
	var node *TreeNode
	if i < len(data) {
		val, err := strconv.Atoi(data[i])
		if err != nil {
			return nil
		}
		node = &TreeNode{Val: val}
		node.Left = binTreeFromSlice(data, 2*i+1)
		node.Right = binTreeFromSlice(data, 2*i+2)
	}
	return node
}

func TreesEqual(t1 *TreeNode, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 != nil && t2 == nil {
		return false
	}
	if t1 == nil && t2 != nil {
		return false
	}
	if t1 != nil && t2 != nil {
		if t1.Val != t2.Val {
			return false
		}
	}
	return TreesEqual(t1.Left, t2.Left) && TreesEqual(t1.Right, t2.Right)
}
