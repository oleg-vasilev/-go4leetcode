package lc112

// 112. Path Sum

// Given the root of a binary tree and an integer targetSum,
// return true if the tree has a root-to-leaf path such that adding
// up all the values along the path equals targetSum.
//
// A leaf is a node with no children.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum-root.Val == 0
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}
