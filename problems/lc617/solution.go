// Package lc617 - Merge binary trees
package lc617

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	newNode := TreeNode{}
	if root1 != nil {
		newNode.Val += root1.Val
	}
	if root2 != nil {
		newNode.Val += root2.Val
	}
	if root1 != nil && root2 != nil {
		newNode.Left = mergeTrees(root1.Left, root2.Left)
		newNode.Right = mergeTrees(root1.Right, root2.Right)
	} else if root1 != nil && root2 == nil {
		newNode.Left = mergeTrees(root1.Left, nil)
		newNode.Right = mergeTrees(root1.Right, nil)
	} else if root1 == nil && root2 != nil {
		newNode.Left = mergeTrees(nil, root2.Left)
		newNode.Right = mergeTrees(nil, root2.Right)
	}
	return &newNode
}
