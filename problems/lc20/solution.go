package lc20

// 20. Valid Parentheses

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

func isValid(s string) bool {
	var stack = Stack[rune]{}
	for _, r := range []rune(s) {
		switch r {
		case '(', '{', '[':
			stack.Push(r)
		case ')', '}', ']':
			if topItem, ok := stack.Pop(); ok && matches(topItem, r) {
				continue
			} else {
				return false
			}
		}
	}
	return stack.Size() == 0
}

func matches(open, close rune) bool {
	return (open == '(' && close == ')') ||
		(open == '{' && close == '}') ||
		(open == '[' && close == ']')
}

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		return *new(T), false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) Top() T {
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}
