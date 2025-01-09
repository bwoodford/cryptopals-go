package eight

import (
	"bufio"
	"os"
	"testing"
)

func TestSetOneChallengeSeven(t *testing.T) {

	file, err := os.Open("set-one-challenge-eight.txt")
	if err != nil {
		t.Fatalf("unable to open encrypted file for testing")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := make([]string, 0)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	have, err := Solve(input)
	if err != nil {
		t.Fatalf("%v", err)
	}

	want := "d880619740a8a19b7840a8a31c810a3d08649af70dc06f4fd5d2d69c744cd283e2dd052f6b641dbf9d11b0348542bb5708649af70dc06f4fd5d2d69c744cd2839475c9dfdbc1d46597949d9c7e82bf5a08649af70dc06f4fd5d2d69c744cd28397a93eab8d6aecd566489154789a6b0308649af70dc06f4fd5d2d69c744cd283d403180c98c8f6db1f2a3f9c4040deb0ab51b29933f2c123c58386b06fba186a"

	if have != want {
		t.Fatalf("Solve(%v...) = '%v', '%v', want '%v'", input[0:10], have, err, want)
	}
}
