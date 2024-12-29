package four

import (
	"cryptopals-go/operations"
	"errors"
	"strings"
)

func Solve(input string) (string, error) {

	lines := strings.Split(input, "\n")
	if len(lines) <= 0 {
		return "", errors.New("unable to split input into lines")
	}

	decBytes, err := operations.HexToBytes(lines[0])
	if err != nil {
		return "", err
	}

	// Score for each iteration of decryption
	iterScores := make([]string, 256)
	// Top scores for each iteration
	topScores := make([]string, 0, 5*len(lines))

	var length int
	var keyExpand []byte

	for _, encValue := range lines {

		decBytes, err = operations.HexToBytes(encValue)
		if err != nil {
			return "", err
		}

		length = len(decBytes)
		keyExpand = make([]byte, length)

		for i := range 256 {

			for j := range length {
				keyExpand[j] = byte(i)
			}

			decryptBytes, err := operations.XOR(decBytes, keyExpand)
			if err != nil {
				return "", nil
			}

			iterScores[i] = operations.ToASCII(decryptBytes)
		}

		topScores = append(topScores, operations.Score(iterScores)[0:5]...)
	}

	return operations.Score(topScores)[0], nil
}
