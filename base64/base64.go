package base64

const padding = '='

var fetchBytes = map[rune]byte{
	'A': 0b000000,
	'B': 0b000001,
	'C': 0b000010,
	'D': 0b000011,
	'E': 0b000100,
	'F': 0b000101,
	'G': 0b000110,
	'H': 0b000111,
	'I': 0b001000,
	'J': 0b001001,
	'K': 0b001010,
	'L': 0b001011,
	'M': 0b001100,
	'N': 0b001101,
	'O': 0b001110,
	'P': 0b001111,

	'Q': 0b010000,
	'R': 0b010001,
	'S': 0b010010,
	'T': 0b010011,
	'U': 0b010100,
	'V': 0b010101,
	'W': 0b010110,
	'X': 0b010111,
	'Y': 0b011000,
	'Z': 0b011001,
	'a': 0b011010,
	'b': 0b011011,
	'c': 0b011100,
	'd': 0b011101,
	'e': 0b011110,
	'f': 0b011111,

	'g': 0b100000,
	'h': 0b100001,
	'i': 0b100010,
	'j': 0b100011,
	'k': 0b100100,
	'l': 0b100101,
	'm': 0b100110,
	'n': 0b100111,
	'o': 0b101000,
	'p': 0b101001,
	'q': 0b101010,
	'r': 0b101011,
	's': 0b101100,
	't': 0b101101,
	'u': 0b101110,
	'v': 0b101111,

	'w': 0b110000,
	'x': 0b110001,
	'y': 0b110010,
	'z': 0b110011,
	'0': 0b110100,
	'1': 0b110101,
	'2': 0b110110,
	'3': 0b110111,
	'4': 0b111000,
	'5': 0b111001,
	'6': 0b111010,
	'7': 0b111011,
	'8': 0b111100,
	'9': 0b111101,
	'+': 0b111110,
	'/': 0b111111,
}

var fetchb64 = map[uint8]rune{
	0b000000: 'A',
	0b000001: 'B',
	0b000010: 'C',
	0b000011: 'D',
	0b000100: 'E',
	0b000101: 'F',
	0b000110: 'G',
	0b000111: 'H',
	0b001000: 'I',
	0b001001: 'J',
	0b001010: 'K',
	0b001011: 'L',
	0b001100: 'M',
	0b001101: 'N',
	0b001110: 'O',
	0b001111: 'P',

	0b010000: 'Q',
	0b010001: 'R',
	0b010010: 'S',
	0b010011: 'T',
	0b010100: 'U',
	0b010101: 'V',
	0b010110: 'W',
	0b010111: 'X',
	0b011000: 'Y',
	0b011001: 'Z',
	0b011010: 'a',
	0b011011: 'b',
	0b011100: 'c',
	0b011101: 'd',
	0b011110: 'e',
	0b011111: 'f',

	0b100000: 'g',
	0b100001: 'h',
	0b100010: 'i',
	0b100011: 'j',
	0b100100: 'k',
	0b100101: 'l',
	0b100110: 'm',
	0b100111: 'n',
	0b101000: 'o',
	0b101001: 'p',
	0b101010: 'q',
	0b101011: 'r',
	0b101100: 's',
	0b101101: 't',
	0b101110: 'u',
	0b101111: 'v',

	0b110000: 'w',
	0b110001: 'x',
	0b110010: 'y',
	0b110011: 'z',
	0b110100: '0',
	0b110101: '1',
	0b110110: '2',
	0b110111: '3',
	0b111000: '4',
	0b111001: '5',
	0b111010: '6',
	0b111011: '7',
	0b111100: '8',
	0b111101: '9',
	0b111110: '+',
	0b111111: '/',
}

// Encode takes a string containing unicode characters and outputs the base64 representation.
func Encode(unicode string) string {

	if len(unicode) == 0 {
		return ""
	}

	encoded := make([]rune, 0, len(unicode))

	mask := uint8(0b1000_0000)
	buffer := uint8(0)
	baseIndex := 0

	uniBytes := []byte(unicode)

	for _, uni := range uniBytes {

		mask = 0b1000_0000

		for j := range 8 {

			buffer = buffer << 1
			// Put the acquired bit to the rightmost bit in the buffer
			buffer |= ((uni & mask) >> ((8 - j) - 1))
			mask = mask >> 1
			baseIndex += 1

			// Get the base64 value when we've filled the buffer with 6 bits.
			if baseIndex == 6 {
				encoded = append(encoded, fetchb64[buffer])
				buffer = 0
				baseIndex = 0
			}
		}
	}

	// Add the last bits in the buffer.
	if baseIndex > 0 {
		remainder := 6 - baseIndex
		// Shift the buffer over the remaining amount.
		encoded = append(encoded, fetchb64[buffer<<remainder])

		// Add artificial padding to the output.
		switch remainder {
		case 2:
			encoded = append(encoded, padding)
		case 4:
			encoded = append(encoded, padding)
			encoded = append(encoded, padding)
		}
	}

	return string(encoded)

}

// Decode takes a string containing base64 characters and outputs the unicode representation.
func Decode(b64Chars string) []byte {

	var decoded []byte

	if len(b64Chars) == 0 {
		return decoded
	}

	decoded = make([]byte, 0, len(b64Chars))
	mask := uint8(0b0010_0000)
	buffer := uint8(0)
	b64Byte := uint8(0)
	index := 0

	for _, b64Char := range b64Chars {
		if b64Char == padding {
			continue
		}

		mask = 0b0010_0000
		b64Byte = fetchBytes[b64Char]

		for j := range 6 {

			buffer = buffer << 1
			buffer |= ((b64Byte & mask) >> ((6 - j) - 1))
			mask = mask >> 1
			index += 1

			if index == 8 {
				decoded = append(decoded, buffer)
				buffer = 0
				index = 0
			}
		}
	}

	return decoded
}
