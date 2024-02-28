package lc101

// 101. Symmetric Tree

// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return symmetric(root.Left, root.Right)
}

func symmetric(left *TreeNode, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	if left.Val != right.Val {
		return false
	}
	// trick here is to correctly choose "inner" and "outer" nodes to compare
	return symmetric(left.Left, right.Right) && symmetric(left.Right, right.Left)
}

func isSymmetricIterative(root *TreeNode) bool {
	if root == nil {
		return true
	}
	// prepare stack with items in left-to-right order
	var stack []*TreeNode
	stack = append(stack, root.Left, root.Right)

	// while stack isn't empty
	for len(stack) > 0 {
		// pop items
		l, r := stack[0], stack[1]
		stack = stack[2:]
		// compare
		if l == nil && r == nil {
			// both are nils - ok
			continue
		} else if (l == nil && r != nil) || l != nil && r == nil {
			// one is nil but other isn't - error
			return false
		} else if l.Val != r.Val {
			// values aren't same - error
			return false
		}
		// append to stack child elements in correct order (tricky without actually drawing a tree)
		// knowing we will compare them in pairs
		stack = append(stack, l.Left, r.Right, l.Right, r.Left)
	}
	return true
}
