package two

import "cryptopals-go/operations"

func Solve(inputOne string, inputTwo string) (string, error) {
	bytesOne, err := operations.HexToBytes("1c0111001f010100061a024b53535009181c")
	if err != nil {
		return "", err
	}

	bytesTwo, err := operations.HexToBytes("686974207468652062756c6c277320657965")
	if err != nil {
		return "", err
	}

	operation, err := operations.XOR(bytesOne, bytesTwo)
	if err != nil {
		return "", err
	}

	return operations.HexToString(operation)
}
