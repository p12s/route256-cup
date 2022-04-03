package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindrom(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name        string
		input       int
		isPalindrom bool
	}{
		{"S", 10, false},
		{"S", 12, false},
		{"S", 656, true},
		{"M", 5555, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.isPalindrom, IsPalindrom(tt.input))
		})
	}
}
