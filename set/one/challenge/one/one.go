package one

import (
	"cryptopals-go/base64"
	"cryptopals-go/operations"
)

func Solve(input string) (string, error) {
	hexBytes, err := operations.HexToBytes(input)
	if err != nil {
		return "", err
	}
	encoding := base64.Encode(string(hexBytes))
	return encoding, err
}
