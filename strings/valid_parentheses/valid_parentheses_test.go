/*
Problem:
- Given a list of parentheses, determine if it is valid.

Example:
- Input: []string{"(", ")", "[", "]", "{", "}"}
  Output: true
- Input: []string{"(", "[", ")", "]"}
  Output: false

Approach:
- Use a stack push the opening parenthesis and pop the last inserted one when
  we encounter a closing parenthesis and check if it is a valid match.

Cost:
- O(n) time, O(n) space.
*/

package main

import (
	"sync"
	"testing"

	"github.com/ferralucho/go-utils-algorithms/datastructures"
	"github.com/stretchr/testify/assert"
)

var (
	m    map[string]string
	once sync.Once
)

func getMap() map[string]string {
	once.Do(func() {
		m = map[string]string{
			"(": ")",
			"{": "}",
			"[": "]",
		}
	})
	return m
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		in       []string
		expected bool
	}{
		{[]string{}, true},
		{[]string{"(", ")"}, true},
		{[]string{"(", ")", "}"}, false},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, IsValid(tt.in))
	}
}

func IsValid(s []string) bool {
	if len(s) == 0 {
		return true
	}

	strMap := getMap()

	stack := datastructures.NewStack()

	for _, p := range s {
		if _, ok := strMap[p]; ok {
			stack.Push(p)
		} else if stack.Size() == 0 || strMap[stack.Pop()] != p {
			return false
		}
	}

	return stack.Size() == 0
}
