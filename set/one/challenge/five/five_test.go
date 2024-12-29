package five

import "testing"

func TestSetOneChallengeFive(t *testing.T) {
	
	message := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	key := []rune{'I', 'C', 'E'}

	want := "0B3637272A2B2E63622C2E69692A23693A2A3C6324202D623D63343C2A26226324272765272A282B2F20430A652E2C652A3124333A653E2B2027630C692B20283165286326302E27282F"
	have, err := Solve(message, key)

	if have != want {
		t.Fatalf("Solve(%v, %q) = %v, %v, want %v", message, key, have, err, want)
	}
}
