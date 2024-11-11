package main

import (
	"testing"
)

// TestBase64 tests if a string of characters is correctly converted to a base64 encoded string.
func TestBase64(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := base64(tt.input)
			if got != tt.expected {
				t.Errorf("base64('%v') != '%v'", tt.input, tt.expected)
			}
		})
	}
}
