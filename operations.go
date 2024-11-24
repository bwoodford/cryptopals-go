package main

import (
	"fmt"
	"unicode"
)

var fetchHexByte = map[rune]byte{
	'0': 0b0000,
	'1': 0b0001,
	'2': 0b0010,
	'3': 0b0011,
	'4': 0b0100,
	'5': 0b0101,
	'6': 0b0110,
	'7': 0b0111,
	'8': 0b1000,
	'9': 0b1001,
	'A': 0b1010,
	'B': 0b1011,
	'C': 0b1100,
	'D': 0b1101,
	'E': 0b1110,
	'F': 0b1111,
}

var fetchHexRune = map[byte]rune{
	0b0000: '0',
	0b0001: '1',
	0b0010: '2',
	0b0011: '3',
	0b0100: '4',
	0b0101: '5',
	0b0110: '6',
	0b0111: '7',
	0b1000: '8',
	0b1001: '9',
	0b1010: 'A',
	0b1011: 'B',
	0b1100: 'C',
	0b1101: 'D',
	0b1110: 'E',
	0b1111: 'F',
}

func XOR(a []byte, b []byte) ([]byte, error) {

	if len(a) != len(b) {
		return nil, fmt.Errorf("input buffers must be the same length")
	}

	result := make([]byte, 0, len(a))

	for i := 0; i < len(a); i++ {
		result = append(result, a[i]^b[i])
	}

	return result, nil
}

func ToBytes(hexString string) ([]byte, error) {

	hexRunes := []rune(hexString)
	length := len(hexRunes)

	if length%2 != 0 {
		return nil, fmt.Errorf("hex string length must be an even number")
	}

	checkMap := func(hexChar rune) (byte, error) {
		hex, ok := fetchHexByte[unicode.ToUpper(hexChar)]
		if !ok {
			return hex, fmt.Errorf("value not found in hex map: %v", hexChar)
		}
		return hex, nil
	}

	result := make([]byte, 0, length)

	first := byte(0)
	second := byte(0)
	var err error

	for i := 0; i < len(hexRunes); i += 2 {
		first, err = checkMap(hexRunes[i])
		if err != nil {
			return nil, err
		}
		second, err = checkMap(hexRunes[i+1])
		if err != nil {
			return nil, err
		}

		result = append(result, (first<<4)|second)
	}

	return result, nil
}

func ToHex(input []byte) (string, error) {

	if len(input)%2 != 0 {
		return "", fmt.Errorf("byte slice length must be an even number")
	}

	checkMap := func(b byte) (rune, error) {
		hex, ok := fetchHexRune[b]
		if !ok {
			return hex, fmt.Errorf("value not found in hex map: %v", b)
		}
		return hex, nil
	}

	result := ""
	var first rune
	var second rune
	var err error

	for _, value := range input {
		first, err = checkMap((value & 0xF0) >> 4)
		if err != nil {
			return result, err
		}
		second, err = checkMap((value & 0x0F))
		if err != nil {
			return result, err
		}

		result += string(first) + string(second)
	}

	return result, nil
}
