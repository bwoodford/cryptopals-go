package seven

import (
	"os"
	"strings"
	"testing"
)

/*
func WriteComparison(have string, want string) {

	haveFile, err := os.Create("have.txt")
	if err != nil {
		panic(err)
	}
	defer haveFile.Close()

	wantFile, err := os.Create("want.txt")
	if err != nil {
		panic(err)
	}
	defer wantFile.Close()


	_, err = haveFile.WriteString(have)
	if err != nil {
		panic(err)
	}

	_, err = wantFile.WriteString(want)
	if err != nil {
		panic(err)
	}
}
*/

func TestSetOneChallengeSeven(t *testing.T) {

	encryptB64, err := os.ReadFile("set-one-challenge-seven-encrypted.txt")
	if err != nil {
		t.Fatalf("unable to open encrypted file for testing")
	}

	input := string(encryptB64)

	have, err := Solve(input)
	if err != nil {
		t.Fatalf("%v", err)
	}
	have = strings.TrimSpace(have)

	decryptedText, err := os.ReadFile("set-one-challenge-seven-decrypted.txt")
	if err != nil {
		t.Fatalf("unable to open decrypted file for testing")
	}
	want := strings.TrimSpace(string(decryptedText))

	if have != want {
		t.Fatalf("Solve(%v...) = '%v', '%v', want '%v'", input[0:10], have, err, want)
	}
}
