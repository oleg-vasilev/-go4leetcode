package lc155

// 155. Min Stack

// Design a stack that supports push, pop, top,
// and retrieving the minimum element in constant time.
//
// Implement the MinStack class:
//
// MinStack() initializes the stack object.
// void push(int val) pushes the element val onto the stack.
// void pop() removes the element on the top of the stack.
// int top() gets the top element of the stack.
// int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.
//

type MinStack struct {
	minIndexes []int
	items      []int
}

func Constructor() MinStack {
	return MinStack{
		items:      make([]int, 0),
		minIndexes: make([]int, 0),
	}
}

func (ms *MinStack) Push(val int) {
	ms.items = append(ms.items, val)
	newItemIdx := len(ms.items) - 1 // last item
	if len(ms.minIndexes) == 0 {
		// currently no items and no min indexes - our item is min
		ms.minIndexes = append(ms.minIndexes, newItemIdx)
		return
	}
	// we already have some items stored - compare min elements
	lastMinIdx := ms.minIndexes[len(ms.minIndexes)-1]
	if val < ms.items[lastMinIdx] {
		ms.minIndexes = append(ms.minIndexes, newItemIdx)
	} else {
		ms.minIndexes = append(ms.minIndexes, lastMinIdx)
	}
}

func (ms *MinStack) Pop() {
	// remove last item
	ms.items = ms.items[:len(ms.items)-1]
	// remove last known min for this state
	ms.minIndexes = ms.minIndexes[:len(ms.minIndexes)-1]
}

func (ms *MinStack) Top() int {
	return ms.items[len(ms.items)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.items[ms.minIndexes[len(ms.minIndexes)-1]]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
