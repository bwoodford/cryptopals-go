package two

import "testing"

func TestSetOneChallengeTwo(t *testing.T) {

	inputOne := "1c0111001f010100061a024b53535009181c"
	inputTwo := "686974207468652062756c6c277320657965"
	want := "746865206B696420646F6E277420706C6179"
	have, err := Solve(inputOne, inputTwo)

	if have != want {
		t.Fatalf("Solve(%v, %v) = %v, %v, want %v", inputOne, inputTwo, have, err, want)
	}
}
