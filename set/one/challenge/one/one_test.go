package one

import "testing"

func TestSetOneChallengeOne(t *testing.T) {

	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	want := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	have, err := Solve(input)

	if have != want {
		t.Fatalf("Solve(%v) = %v, %v, want %v", input, have, err, want)
	}

}
