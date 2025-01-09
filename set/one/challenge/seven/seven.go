package seven

import (
	"crypto/aes"
	"encoding/base64"
)

func Solve(encryptB64 string) (string, error) {

	encryptText, err := base64.StdEncoding.DecodeString(encryptB64)
	if err != nil {
		return "", err
	}

	encryptBytes := []byte(encryptText)
	length := len(encryptBytes)

	key := "YELLOW SUBMARINE"
	cBlock, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	decryptBytes := make([]byte, length)

	for i := 0; i < length; i+=aes.BlockSize {
		cBlock.Decrypt(decryptBytes[i:i+aes.BlockSize], encryptBytes[i:i+aes.BlockSize])
	}

	return string(decryptBytes), nil
}
