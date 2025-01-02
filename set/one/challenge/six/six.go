package six

import (
	"cmp"
	"cryptopals-go/operations"
	stdb64 "encoding/base64"
	"fmt"
	"slices"
)


func findKeySizes(encryptBytes []byte, numKeys int) ([]int, error) {

	if numKeys <= 0 {
		return []int{}, nil
	}

	const MaxKeySize = 40

	if numKeys > 40 {
		return []int{}, fmt.Errorf("numKeys can not be greater than %v", MaxKeySize)
	}

	type keyTest struct {
		size int
		distance float64
	}

	keyTests := make([]keyTest, 0, MaxKeySize)
	keySizes := make([]int, 0, numKeys)

	var distance float64
	var err error
	var tmpD int

	var bytesOne []byte
	var bytesTwo []byte
	var bytesThree []byte
	var bytesFour []byte

	// We'll just add on 2 more key iterations for semantics
	for i := 2; i < MaxKeySize+2; i++ {
		distance = 0

		bytesOne = encryptBytes[:i]
		bytesTwo = encryptBytes[i:i*2]
		bytesThree = encryptBytes[i*2:i*3]
		bytesFour = encryptBytes[i*3:i*4]

		tmpD, err = operations.Distance(bytesOne, bytesTwo)
		if err != nil {
			return keySizes, err
		}
		distance += float64(tmpD)

		tmpD, err = operations.Distance(bytesOne, bytesThree)
		if err != nil {
			return keySizes, err
		}
		distance += float64(tmpD)

		tmpD, err = operations.Distance(bytesOne, bytesFour)
		if err != nil {
			return keySizes, err
		}
		distance += float64(tmpD)

		tmpD, err = operations.Distance(bytesTwo, bytesThree)
		if err != nil {
			return keySizes, err
		}
		distance += float64(tmpD)

		tmpD, err = operations.Distance(bytesTwo, bytesFour)
		if err != nil {
			return keySizes, err
		}
		distance += float64(tmpD)

		tmpD, err = operations.Distance(bytesThree, bytesFour)
		if err != nil {
			return keySizes, err
		}
		distance += float64(tmpD)

		keyTests = append(keyTests, keyTest{distance: ((distance/6)/float64(i)), size: i})
	}

	slices.SortFunc(keyTests, func(a keyTest, b keyTest) int {
		return cmp.Compare(a.distance, b.distance)
	})

	for i := range numKeys {
		keySizes = append(keySizes, keyTests[i].size)
	}

	return keySizes, nil
}

func Solve(encryptB64 string) (string, error){

	decryptText := ""

	encryptBytes, err := stdb64.StdEncoding.DecodeString(encryptB64)
	if err != nil {
		return "", err
	}

	encryptLength := len(encryptBytes)

	keySizes, err := findKeySizes(encryptBytes, 1)
	if err != nil {
		return "", err
	}

	keySize := keySizes[0]

	transposed := make([][]byte, keySize)

	for i := 0; i < keySize; i++ {
		for j:= i; j < encryptLength; j += keySize {
			transposed[i] = append(transposed[i], encryptBytes[j])
		}
	}

	type keyScore struct {
		key byte
		score float64
	}

	key := make([]byte, keySize)
	
	for i, block := range transposed {

		keyScores := make([]keyScore, 0, 256)
		tmpKey := make([]byte, len(block))

		for j := range 256 {

			for k := range len(block) {
				tmpKey[k] = byte(j)
			}

			decryptBytes, err := operations.XOR(block, tmpKey)
			if err != nil {
				return "", err
			}
			keyScores = append(keyScores, keyScore{ key: byte(j), score: operations.ChiSquaredScore(operations.ToASCII(decryptBytes))})
		}

		slices.SortFunc(keyScores, func (a keyScore, b keyScore) int {
			return cmp.Compare(a.score, b.score)
		})

		key[i] = keyScores[0].key
	}

	for i := 0; i < len(encryptBytes); i++ {
		decryptByte, err := operations.XOR([]byte{key[i%keySize]}, []byte{encryptBytes[i]})
		if err != nil {
			return "", err
		}
		decryptText += string(decryptByte)
	}

	return decryptText, nil
}
