package operations

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"unicode"
)


// Taken from https://en.wikipedia.org/wiki/Letter_frequency
var characterFrequency = map[rune]float64{
	'A': .082,
	'B': .015,
	'C': .028,
	'D': .043,
	'E': .127,
	'F': .022,
	'G': .020,
	'H': .061,
	'I': .070,
	'J': .002,
	'K': .008,
	'L': .040,
	'M': .024,
	'N': .067,
	'O': .075,
	'P': .019,
	'Q': .001,
	'R': .060,
	'S': .063,
	'T': .091,
	'U': .028,
	'V': .010,
	'W': .024,
	'X': .002,
	'Y': .020,
	'Z': .001,
}

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

func HexToBytes(hexString string) ([]byte, error) {

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

func HexToString(hexBytes []byte) (string, error) {

	if len(hexBytes)%2 != 0 {
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

	for _, b := range hexBytes {
		first, err = checkMap((b & 0xF0) >> 4)
		if err != nil {
			return result, err
		}
		second, err = checkMap((b & 0x0F))
		if err != nil {
			return result, err
		}

		result += string(first) + string(second)
	}

	return result, nil
}

func ToASCII(input []byte) string {
	return string(input)
}

func Score(input []string) []string {

	calc := func(decrypt string) int {
		score := 0

		var upper rune
		var ok bool

		for _, char := range decrypt {
			upper = unicode.ToUpper(char)
			_, ok = characterFrequency[upper]
			if ok  || upper == ' '{
				score += 1
			}
		}

		return score
	}

	slices.SortFunc(input, func(a, b string) int {
		return -cmp.Compare(calc(a), calc(b))
	})

	return input
}

func Distance(a []byte, b []byte) (int, error) {

	if len(a) != len(b) {
		return 0, errors.New("inputs must be the same length")
	}

	distance := 0

	diffBytes, err := XOR(a, b)
	if err != nil {
		return 0, err
	}

	var mask byte

	for i := 0; i < len(diffBytes); i++ {

		mask = 0b1000_0000
		for j := range byte(8) {
			distance += int((diffBytes[i] & mask) >> (7 - j))
			mask = mask >> 1
		}
	}

	return distance, nil
}

func ChiSquaredScore(input string) float64 {

	length := len(input)
	score := float64(0)

	characters := make(map[rune]int, 26)
	ignore := 0
	for _, r := range input {
		if int(r) == 32 || int(r) == 13 || int(r) == 9 {
			ignore += 1
		} else {
			// Maybe add some ignores for characters we actually care about like spaces, carriage returns...etc
			characters[unicode.ToUpper(r)] += 1
		}
	}

	length = length - ignore

	var freq float64
	var ok bool
	for r, observed := range characters {
		freq, ok = characterFrequency[r]
		if !ok {
			freq = 0.001
		}
		expected := float64(length) * freq
		diff := (float64(observed)-expected)
		score += ((diff * diff) / expected)
	}

	return score
}
