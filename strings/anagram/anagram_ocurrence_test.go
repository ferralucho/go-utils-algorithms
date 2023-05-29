/*
Problem:
- Given an array of strings, group anagrams together.
- All inputs will be in lowercase.
- The order of your output does not matter.

Example:
- Input: []string{"eat", "tea", "tan", "ate", "nat", "bat"}
  Output: [][]string{
	  []string{"ate", "eat", "tea"},
	  []string{"nat","tan"},
	  []string{"bat"},
  }

Approach:
- Two strings are anagrams if and only if their character counts
  (respective number of occurrences of each character) are the same.

Cost:
- O(n) time, O(n) space.
*/

package main

import (
	"reflect"
	"sort"
	"strings"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expected := [][]string{
		{"eat", "tea", "ate"},
		{"tan", "nat"},
		{"bat"},
	}

	result := groupAnagrams(strs)

	if len(result) != len(expected) {
		t.Errorf("Unexpected group count. Expected: %d, Got: %d", len(expected), len(result))
	}

	for _, group := range expected {
		if !containsGroup(result, group) {
			t.Errorf("Group not found: %v", group)
		}
	}
}

func containsGroup(groups [][]string, group []string) bool {
	for _, g := range groups {
		if reflect.DeepEqual(g, group) {
			return true
		}
	}
	return false
}

func sortString(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)

	for _, str := range strs {
		sorted := sortString(str)
		anagrams[sorted] = append(anagrams[sorted], str)
	}

	groups := make([][]string, 0, len(anagrams))

	for _, v := range anagrams {
		groups = append(groups, v)
	}
	return groups
}
