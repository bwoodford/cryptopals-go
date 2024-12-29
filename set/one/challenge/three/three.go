package three

import "cryptopals-go/operations"

func Solve(input string) (string, error){

	bytesValue, err := operations.HexToBytes(input)
	if err != nil {
		return "", err
	}

	length := len(bytesValue)
	keyExpand := make([]byte, length)
	decryptValues := make([]string, 256)

	for i := range 256 {

		for j := range length {
			keyExpand[j] = byte(i)
		}

		decryptBytes, err := operations.XOR(bytesValue, keyExpand)
		if err != nil {
			return "", err
		}

		decryptValues[i] = operations.ToASCII(decryptBytes)
	}

	return operations.Score(decryptValues)[0], nil
}
