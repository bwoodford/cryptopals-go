package six

import (
	"os"
	"testing"
)

func TestSetOneChallengeSix(t *testing.T) {

	encryptB64, err := os.ReadFile("set-one-challenge-six-encrypted.txt")
	if err != nil {
		t.Fatalf("unable to open encrypted file for testing")
	}

	input := string(encryptB64)

	have, err := Solve(input)
	if err != nil {
		t.Fatalf("%v", err)
	}

	decryptedText, err := os.ReadFile("set-one-challenge-six-decrypted.txt")
	if err != nil {
		t.Fatalf("unable to open decrypted file for testing")
	}

	want := string(decryptedText)

	if have != want {
		t.Fatalf("Solve(%v...) = '%v', '%v', want '%v'", input[0:10], have, err, want)
	}

}
