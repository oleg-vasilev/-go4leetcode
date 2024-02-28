package lc100

// 100. Same Tree

// Given the roots of two binary trees p and q,
// write a function to check if they are the same or not.
//
// Two binary trees are considered the same if they are structurally identical,
// and the nodes have the same value.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// straightforward recursive dfs
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p == q { // handles both nil case and if somehow p and q points to same object
		return true
	}
	// values are same - check childs
	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// iterative
func isSameTree2(p *TreeNode, q *TreeNode) bool {
	if p == nil && q != nil {
		return false
	}
	if p != nil && q == nil {
		return false
	}
	if p == q { // handles both nil case and if somehow p and q points to same object
		return true
	}

	// Stacks for a depth-first search
	pStack, qStack := []*TreeNode{p}, []*TreeNode{q}

	// Pointers to each tree's current node
	var pCurr, qCurr *TreeNode

	for len(pStack) > 0 && len(qStack) > 0 {
		// Pop top nodes from each stack
		pCurr, pStack = pStack[len(pStack)-1], pStack[:len(pStack)-1]
		qCurr, qStack = qStack[len(qStack)-1], qStack[:len(qStack)-1]
		// compare values
		if pCurr.Val != qCurr.Val {
			return false
		}
		// process childrens
		// asymmetric cases - trees are not the same by definition
		if (pCurr.Left != nil && qCurr.Left == nil) || // left 1
			(pCurr.Left == nil && qCurr.Left != nil) || // left 2
			(pCurr.Right != nil && qCurr.Right == nil) || // right 1
			(pCurr.Right == nil && qCurr.Right != nil) { // right 2
			return false
		}
		// push to stacks child nodes if both node have ones
		if pCurr.Left != nil && qCurr.Left != nil {
			pStack = append(pStack, pCurr.Left)
			qStack = append(qStack, qCurr.Left)
		}
		if pCurr.Right != nil && qCurr.Right != nil {
			pStack = append(pStack, pCurr.Right)
			qStack = append(qStack, qCurr.Right)
		}
	}
	// trees are the same if at the end both stacks are empty
	return len(pStack) == len(qStack)
}
