package four

import (
	"os"
	"testing"
)

func TestSetOneChallengeFour(t *testing.T) {

	contents, err := os.ReadFile("set-one-challenge-four.txt")
	if err != nil {
		t.Fatalf("unable to open file for testing")
	}

	input := string(contents)
	want := "Now that the party is jumping\n"
	have, err := Solve(input)

	if have != want {
		t.Fatalf("Solve(%v...) = '%v', '%v', want '%v'", input[0:10], have, err, want)
	}

}
