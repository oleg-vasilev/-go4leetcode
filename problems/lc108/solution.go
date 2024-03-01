package lc108

// 108. Convert Sorted Array to Binary Search Tree

// Given an integer array nums where the elements are sorted in ascending
// order, convert it to a height-balanced binary search tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	// base case
	if len(nums) == 0 {
		return nil
	}
	var center = len(nums) / 2
	var node = new(TreeNode)
	node.Val = nums[center]
	if center > 0 {
		node.Left = sortedArrayToBST(nums[:center])
	}
	if center < len(nums)-1 {
		node.Right = sortedArrayToBST(nums[center+1:])
	}
	return node
}
