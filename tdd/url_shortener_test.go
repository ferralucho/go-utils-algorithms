package tdd

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
// Step 1: Write a failing test
func TestURLShortner(t *testing.T) {
	var tests = []struct {
		name, in, expected string
		err                error
	}{
		{"NoScheme", "https://google.com", "1", nil},
		{"NoScheme", "google", "", errors.New("invalid")},
	}

	for _, test := range tests {
		t.Run("Shorten"+test.name, func(t *testing.T) {
			s := &URLShortner{i: 0, store: make(map[string]string)}
		})
	}
}
*/

// Step 3: Write test code to test the behaviour
func TestURLShortner(t *testing.T) {
	var tests = []struct {
		name, in, expected string
		err                error
	}{
		{"NoScheme", "https://google.com", "1", nil},
		{"NoScheme", "google", "", errors.New("invalid")},
	}

	for _, test := range tests {
		t.Run("Shorten"+test.name, func(t *testing.T) {
			s := &URLShortener{i: 0, store: make(map[string]string)}
			resp, err := s.Shorten(test.in)
			assert.Equal(t, test.err, err)
			assert.Equal(t, test.expected, resp)
		})
	}
}
