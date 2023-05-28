/*
Problem:
- Given a string, determine if it is a palindrome, considering only
  alphanumeric characters and ignoring cases.

Example:
- Input: "A man, a plan, a canal: Panama"
  Output: true
- Input: "race a car"
  Output: false

Approach:
- Use two pointers approach that have one point to the start of the string and
  the other point at the end.
- Move them toward each other and compare if they're the same characters while
  skipping non-alphanumeric characters and ignoring cases.

Solution:
- Have a pointer point to the start of the string and the other point at the end.
- Make the string lower case.
- While the start one is less the end one, ignore non-alphanumeric characters
  and move them toward each other.

Cost:
- O(n) time, O(1) space.
*/

package main

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
)

func TestIsPalindrome(t *testing.T) {
	input1 := "A man, a plan, a canal, Panama"
	expected1 := true
	result1 := IsPalindrome(input1)
	if result1 != expected1 {
		t.Errorf("Test case 1 failed. Expected: %v, got: %v", expected1, result1)
	}

	input2 := "Hello, world!"
	expected2 := false
	result2 := IsPalindrome(input2)
	if result2 != expected2 {
		t.Errorf("Test case 2 failed. Expected: %v, got: %v", expected2, result2)
	}

	input3 := ""
	expected3 := true
	result3 := IsPalindrome(input3)
	if result3 != expected3 {
		t.Errorf("Test case 3 failed. Expected: %v, got: %v", expected3, result3)
	}
}

func IsPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	s = strings.ToLower(s)

	for start < end {
		for !unicode.IsLetter(getRune(s, start)) {
			start++
		}
		for !unicode.IsLetter(getRune(s, end)) {
			end--
		}

		if getRune(s, start) != getRune(s, end) {
			return false
		}

		start++
		end--
	}
	return true
}

func getRune(s string, pos int) rune {
	return []rune(s)[pos]
}

func main() {
	sn := "karate"
	s := "anana"

	fmt.Println(sn)
	fmt.Println(s)
}
