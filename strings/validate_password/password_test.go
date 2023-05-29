package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"", "DEBIL"},
		{"rr", "DEBIL"},
		{"12345678", "BUENA"},
		{"aaaabbbb", "BUENA"},
		{"23g", "BUENA"},
		{"123ttY8831", "EXCELENTE"},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.expected, ValidatePassword(tt.in))
	}
}
