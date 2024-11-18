package base64

import (
	"unicode"
)

const padding = "="

var fetchHex = map[rune]uint8{
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

var fetchb64 = map[uint8]string{
	0b000000: "A",
	0b000001: "B",
	0b000010: "C",
	0b000011: "D",
	0b000100: "E",
	0b000101: "F",
	0b000110: "G",
	0b000111: "H",
	0b001000: "I",
	0b001001: "J",
	0b001010: "K",
	0b001011: "L",
	0b001100: "M",
	0b001101: "N",
	0b001110: "O",
	0b001111: "P",

	0b010000: "Q",
	0b010001: "R",
	0b010010: "S",
	0b010011: "T",
	0b010100: "U",
	0b010101: "V",
	0b010110: "W",
	0b010111: "X",
	0b011000: "Y",
	0b011001: "Z",
	0b011010: "a",
	0b011011: "b",
	0b011100: "c",
	0b011101: "d",
	0b011110: "e",
	0b011111: "f",

	0b100000: "g",
	0b100001: "h",
	0b100010: "i",
	0b100011: "j",
	0b100100: "k",
	0b100101: "l",
	0b100110: "m",
	0b100111: "n",
	0b101000: "o",
	0b101001: "p",
	0b101010: "q",
	0b101011: "r",
	0b101100: "s",
	0b101101: "t",
	0b101110: "u",
	0b101111: "v",

	0b110000: "w",
	0b110001: "x",
	0b110010: "y",
	0b110011: "z",
	0b110100: "0",
	0b110101: "1",
	0b110110: "2",
	0b110111: "3",
	0b111000: "4",
	0b111001: "5",
	0b111010: "6",
	0b111011: "7",
	0b111100: "8",
	0b111101: "9",
	0b111110: "+",
	0b111111: "/",
}

// Encode takes a string containing hexadecimal values and outputs the base64 representation.
func Encode(hexChars string) string {

	if len(hexChars) == 0 {
		return ""
	}

	encoded := ""

	mask := uint8(0b1000)
	buffer := uint8(0)
	hex := uint8(0)
	baseIndex := 0

	for _, char := range hexChars {

		mask = uint8(0b1000)
		hex = fetchHex[unicode.ToUpper(char)]

		// Iterate through each hex value (4 bits).
		for j := range 4 {

			buffer = buffer << 1
			// Put the acquired bit to the rightmost bit in the buffer
			buffer |= ((hex & mask) >> ((4 - j) - 1))
			mask = mask >> 1
			baseIndex += 1

			// Get the base64 value when we've filled the buffer with 6 bits.
			if baseIndex == 6 {
				encoded += fetchb64[buffer]
				buffer = 0
				baseIndex = 0
			}
		}
	}

	// Add the last bits in the buffer.
	if baseIndex > 0 {
		remainder := 6 - baseIndex
		// Shift the buffer over the remaining amount.
		encoded += fetchb64[buffer<<remainder]

		// Add artificial padding to the output string.
		switch remainder {
		case 2:
			encoded += padding
		case 4:
			encoded += padding + padding
		}
	}

	return encoded

}
