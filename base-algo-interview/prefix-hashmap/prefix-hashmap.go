/*
HashMap or Trie for Common Substrings
- Use HashMaps or Tries for substring searches and prefix matching.
- They efficiently handle string patterns and reduce redundant checks.
- Example: Find the longest common prefix among multiple strings.
*/

package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}

	return prefix
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	prefix := longestCommonPrefix(strs)
	fmt.Println("Longest Common Prefix:", prefix) // Output: Longest Common Prefix: fl
}
