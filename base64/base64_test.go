package base64

import (
	"testing"
)

// TestEncode tests if a string representing hexadecimal values is correctly converted to a base64 encoded string.
func TestEncode(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
		{"", ""},
		{"4d", "TQ=="},
		{"61", "YQ=="},
		{"7a", "eg=="},
		{"4d61", "TWE="},
		{"7a7a", "eno="},
		{"414243", "QUJD"},
		{"48656c6c6f", "SGVsbG8="},
		{"576f726c64", "V29ybGQ="},
		{"54657374", "VGVzdA=="},
		{"4578616d706c65", "RXhhbXBsZQ=="},
		{"2021", "ICE="},
		{"3f3f", "Pz8="},
		{"2323", "IyM="},
		{"48656c6c6f20576f726c642121", "SGVsbG8gV29ybGQhIQ=="},
		{"54686973206973206120746573742e", "VGhpcyBpcyBhIHRlc3Qu"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := Encode(tt.input)
			if got != tt.want {
				t.Errorf("got: '%v', want: '%v'", got, tt.want)
			}
		})
	}
}
