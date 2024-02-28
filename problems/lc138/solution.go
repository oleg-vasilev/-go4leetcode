package lc138

// 138. Copy List with Random Pointer

// A linked list of length n is given such that each node contains an
// additional random pointer, which could point to any node in the list, or null.
//
// Construct a deep copy of the list. The deep copy should consist of
// exactly n brand new nodes, where each new node has its value set to
// the value of its corresponding original node. Both the next and random
// pointer of the new nodes should point to new nodes in the copied list
// such that the pointers in the original list and copied list represent
// the same list state. None of the pointers in the new list
// should point to nodes in the original list.
//
// For example, if there are two nodes X and Y in the original list,
// where X.random --> Y, then for the corresponding two nodes x and y
// in the copied list, x.random --> y.
//
// Return the head of the copied linked list.
//
// The linked list is represented in the input/output as a list of n nodes.
// Each node is represented as a pair of [val, random_index] where:
//
// val: an integer representing Node.val
// random_index: the index of the node (range from 0 to n-1) that the random
// pointer points to, or null if it does not point to any node.
// Your code will only be given the head of the original linked list.

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// straightforward solution using additional space to map nodes
func copyRandomList(head *Node) *Node {
	var old2new = make(map[*Node]*Node)

	var newHead, newPrev, newPtr *Node
	var oldPtr = head

	for oldPtr != nil {
		// check if we already have matching node for this old node
		if existingNode, ok := old2new[oldPtr]; ok {
			newPtr = existingNode
		} else {
			newPtr = new(Node)
			newPtr.Val = oldPtr.Val
			old2new[oldPtr] = newPtr
		}
		// link new chain together
		if newPrev != nil {
			newPrev.Next = newPtr
		}

		// check link to random node
		if oldPtr.Random != nil {
			// if new chain already has suitable element
			if existingRandomNode, ok := old2new[oldPtr.Random]; ok {
				newPtr.Random = existingRandomNode
			} else {
				// create new node for random link and register it in map
				newPtr.Random = new(Node)
				newPtr.Random.Val = oldPtr.Random.Val
				old2new[oldPtr.Random] = newPtr.Random
			}
		}

		// handle case for first iteration
		if newHead == nil {
			newHead = newPtr
		}
		// do shifts for next iteration
		oldPtr = oldPtr.Next
		newPrev = newPtr
		newPtr = nil
	}
	return newHead
}

// same approach with map but more simple logically
func copyRandomList2(head *Node) *Node {
	var nodesMap = make(map[*Node]*Node)
	var current = head

	// 1st step - register all old nodes in map with their copies
	for current != nil {
		nodesMap[current] = &Node{Val: current.Val}
		current = current.Next
	}
	// 2nd step - iterate over old chain again to setup links between copies
	current = head
	for current != nil {
		node := nodesMap[current]
		node.Next = nodesMap[current.Next]
		node.Random = nodesMap[current.Random]
		current = current.Next
	}

	return nodesMap[head]
}

func copyRandomListRecursive(head *Node) *Node {
	known := make(map[*Node]*Node)
	return copyRecursive(head, known)
}

func copyRecursive(head *Node, known map[*Node]*Node) *Node {
	if head == nil {
		return nil
	}

	if node, ok := known[head]; ok {
		return node
	}

	node := new(Node)
	node.Val = head.Val

	known[head] = node

	node.Next = copyRecursive(head.Next, known)
	node.Random = copyRecursive(head.Random, known)

	return node
}
