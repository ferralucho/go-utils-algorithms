package main

func longestSubstringWithoutRepeatingCharacters(s string) string {
	// map to store the last index at which each character appeared.
	lastSeen := make(map[rune]int)

	start := 0
	maxLength := 0
	maxStart := 0

	for i, c := range s {
		if lastIndex, ok := lastSeen[c]; ok && lastIndex >= start {
			start = lastIndex + 1
		} else {
			length := i - start + 1
			if length > maxLength {
				maxLength = length
				maxStart = start
			}
		}

		lastSeen[c] = i
	}

	return s[maxStart : maxStart+maxLength]
}
