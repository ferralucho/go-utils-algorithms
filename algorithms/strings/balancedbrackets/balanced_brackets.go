package main

import "fmt"

/*
Exercise
Given an expression string exp, write a program to examine
whether the pairs and the orders of “{“, “}”, “(“, “)”, “[“, “]” are correct in exp.

Input: exp = “[()]{}{[()()]()}”
Output: Balanced
Input: exp = “[(])”
Output: Not Balanced
*/

var openPairs = map[rune]bool{
	'{': true,
	'(': true,
	'[': true,
}

var closePairs = map[rune]bool{
	'}': true,
	')': true,
	']': true,
}

var matchPairs = map[rune]rune{
	'}': '{',
	')': '(',
	']': '[',
}

func IsBalancedPairs(s string) bool {
	if s == "" {
		return true
	}

	stack := []rune{}

	for _, c := range s {
		if _, isOpen := openPairs[c]; isOpen {
			stack = append(stack, c)
			continue
		}
		if _, isClosed := closePairs[c]; isClosed {
			if len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == matchPairs[c] {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

func main() {
	result := IsBalancedPairs("[()]{}{[()()]()}")
	if result {
		fmt.Println("Balanced")
	} else {
		fmt.Println("Unbalanced")
	}
}
