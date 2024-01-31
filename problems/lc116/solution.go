// Package lc116 - Populating Next Right Pointers in Each Node
package lc116

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	if root.Left != nil && root.Right != nil {
		root.Left.Next = root.Right
		_ = connect(root.Left)
		if root.Next != nil {
			root.Right.Next = root.Next.Left
		}
		_ = connect(root.Right)
	}
	return root
}
