package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSome(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"S", "123", []string{
			"123",
			"132",
			"213",
			"231",
			"312",
			"321",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tt.expected, Some(tt.input))
		})
	}
}
