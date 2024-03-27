package lc230

// 230. Kth Smallest Element in a BST

// Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of all the values of the nodes in the tree.
//

// Constraints:
//
// The number of nodes in the tree is n.
// 1 <= k <= n <= 104
// 0 <= Node.val <= 104

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	// do dfs inorder traversal and count items
	var counter int = 1
	var stack = make([]*TreeNode, 0)

	var current = root
	for len(stack) > 0 || current != nil {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else {
			// pop from stack
			current = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// visit node
			if counter == k {
				return current.Val
			} else {
				counter++
			}
			// go to right side
			current = current.Right
		}
	}
	return -1
}

func dfs(root *TreeNode, k *int) int {
	if root == nil {
		return 0
	}
	// left child
	if val := dfs(root.Left, k); val > 0 {
		return val
	}
	// decrease counter for current node
	*k -= 1
	// if counter already zero = return current root value
	if *k == 0 {
		return root.Val
	}
	// or go right subtree otherwise
	return dfs(root.Right, k)

}

func kthSmallestRecursive(root *TreeNode, k int) int {
	return dfs(root, &k)
}
