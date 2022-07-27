package main

import (
	"testing"
	"unicode/utf8"
)

func TestReverse(t *testing.T) {
	testCases := []struct{ in, want string }{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}

	for _, tc := range testCases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse %q: want %q", rev, tc.want)
		}

		if !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	}
}
