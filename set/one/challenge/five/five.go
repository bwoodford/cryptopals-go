package five

import "cryptopals-go/operations"

func Solve(message string, key []rune) (string, error) {
	messLength := len(message)
	keyLength := len(key)
	keyExpand := make([]byte, messLength)

	for i := range messLength {
		keyExpand[i] = byte(key[i%keyLength])
	}

	encBytes, err := operations.XOR([]byte(message), keyExpand)
	if err != nil {
		return "", err
	}

	return operations.HexToString(encBytes)
}
