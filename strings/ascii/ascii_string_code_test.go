package main

import (
	"fmt"
	"testing"
)

func TestDecodePassword(t *testing.T) {
	// Test case 1: Sample input
	password := "796115110113721110141108"
	expected := "PrepInsta"
	result := decodePassword(password)
	fmt.Printf("Input: %s\nExpected: %s\nResult: %s\n\n", password, expected, result)

	// Test case 2: Password with spaces
	password = "3210111011697"
	expected = " PrepInsta"
	result = decodePassword(password)
	fmt.Printf("Input: %s\nExpected: %s\nResult: %s\n\n", password, expected, result)

	// Test case 3: Password with lowercase and uppercase characters
	password = "10011411710932111997102115101109"
	expected = "mOT HeR cOde"
	result = decodePassword(password)
	fmt.Printf("Input: %s\nExpected: %s\nResult: %s\n\n", password, expected, result)

	// Test case 4: Password with only spaces
	password = "321032101092032101012132032109107"
	expected = " Pre p In  s ta k"
	result = decodePassword(password)
	fmt.Printf("Input: %s\nExpected: %s\nResult: %s\n\n", password, expected, result)
}

func decodePassword(password string) string {
	reversed := reverseString(password)

	decoded := ""
	i := 0
	n := len(reversed)

	for i < n {
		// Get the current digit or two-digit number
		var current string
		if i+1 < n {
			current = reversed[i : i+2]
		} else {
			current = string(reversed[i])
		}

		ascii := stringToASCII(current)

		decoded += fmt.Sprint(ascii)

		// Move the index based on the number of digits in the current value
		if i+1 < n {
			i += 2
		} else {
			i++
		}
	}

	return decoded
}

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

func stringToASCII(s string) int {
	if len(s) == 2 {
		return int(s[0]-'0')*10 + int(s[1]-'0')
	}
	return int(s[0])
}

func main() {
	password := "796115110113721110141108"
	decoded := decodePassword(password)
	fmt.Println(decoded)
}
