package eight

import (
	"crypto/aes"
	"cryptopals-go/operations"
	"errors"
)

func Solve(input []string) (string, error) {
	var blocks map[string]bool

	for _, line := range input {

		hex, err := operations.HexToBytes(line)
		if err != nil {
			return "", err
		}

		blocks = make(map[string]bool, len(hex)/aes.BlockSize)

		for i := 0; i < len(hex); i += aes.BlockSize {
			_, ok := blocks[string(hex[i:i+aes.BlockSize])]
			if ok {
				return line, nil
			}
			blocks[string(hex[i:i+aes.BlockSize])] = true
		}
	}

	return "", errors.New("No match was found :(")
}
