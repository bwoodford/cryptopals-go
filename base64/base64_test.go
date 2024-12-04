package base64

import (
	"testing"
)

// TestEncode tests if a string with unicdoe values is correctly converted to a base64 encoded string.
func TestEncode(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d", "NDkyNzZkMjA2YjY5NmM2YzY5NmU2NzIwNzk2Zjc1NzIyMDYyNzI2MTY5NmUyMDZjNjk2YjY1MjA2MTIwNzA2ZjY5NzM2ZjZlNmY3NTczMjA2ZDc1NzM2ODcyNmY2ZjZk"},
		{"", ""},
		{"4d", "NGQ="},
		{"61", "NjE="},
		{"7a", "N2E="},
		{"4d61", "NGQ2MQ=="},
		{"7a7a", "N2E3YQ=="},
		{"414243", "NDE0MjQz"},
		{"48656c6c6f", "NDg2NTZjNmM2Zg=="},
		{"576f726c64", "NTc2ZjcyNmM2NA=="},
		{"54657374", "NTQ2NTczNzQ="},
		{"4578616d706c65", "NDU3ODYxNmQ3MDZjNjU="},
		{"2021", "MjAyMQ=="},
		{"3f3f", "M2YzZg=="},
		{"2323", "MjMyMw=="},
		{"48656c6c6f20576f726c642121", "NDg2NTZjNmM2ZjIwNTc2ZjcyNmM2NDIxMjE="},
		{"54686973206973206120746573742e", "NTQ2ODY5NzMyMDY5NzMyMDYxMjA3NDY1NzM3NDJl"},
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
