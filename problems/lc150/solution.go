package lc150

import (
	"strconv"
)

// 150. Evaluate Reverse Polish Notation

// You are given an array of strings tokens that represents an
// arithmetic expression in a Reverse Polish Notation.
//
// Evaluate the expression.
// Return an integer that represents the value of the expression.
//
// Note that:
//
// The valid operators are '+', '-', '*', and '/'.
// Each operand may be an integer or another expression.
// The division between two integers always truncates toward zero.
// There will not be any division by zero.
// The input represents a valid arithmetic expression in a reverse polish notation.
// The answer and all the intermediate calculations can be represented in a 32-bit integer.

func evalRPN(tokens []string) int {

	var stack = make([]int, 0, len(tokens))

	var res int

	for _, token := range tokens {
		switch token {

		case multiply:
			res = stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		case divide:
			res = stack[len(stack)-2] / stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		case subtract:
			res = stack[len(stack)-2] - stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		case sum:
			res = stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, res)
		default:
			// number
			value, _ := strconv.Atoi(token)
			stack = append(stack, value)
		}

	}
	return stack[len(stack)-1]
}

const (
	multiply = "*"
	divide   = "/"
	subtract = "-"
	sum      = "+"
)
