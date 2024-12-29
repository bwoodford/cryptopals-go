package operations

import (
	"bytes"
	"fmt"
	"testing"
)

func TestToBytes(t *testing.T) {
	tests := []struct {
		input string
		want  []byte
	}{
		{"AB", []byte{0b1010_1011}},
		{"1c0111001f010100061a024b53535009181c", []byte{0b0001_1100, 0b0000_0001, 0b0001_0001, 0b0000_0000, 0b0001_1111, 0b0000_0001, 0b0000_0001, 0b0000_0000, 0b0000_0110, 0b0001_1010, 0b0000_0010, 0b0100_1011, 0b0101_0011, 0b0101_0011, 0b0101_0000, 0b0000_1001, 0b0001_1000, 0b0001_1100}},
		{"686974207468652062756c6c277320657965", []byte{0b0110_1000, 0b0110_1001, 0b0111_0100, 0b0010_0000, 0b0111_0100, 0b0110_1000, 0b0110_0101, 0b0010_0000, 0b0110_0010, 0b0111_0101, 0b0110_1100, 0b0110_1100, 0b0010_0111, 0b0111_0011, 0b0010_0000, 0b0110_0101, 0b0111_1001, 0b0110_0101}},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got, err := HexToBytes(tt.input)
			if !bytes.Equal(got, tt.want) || err != nil {
				t.Errorf("got: '%v', want: '%v', err: %v", got, tt.want, err)
			}
		})
	}
}

func TestXOR(t *testing.T) {
	type inputStruct struct {
		a []byte
		b []byte
	}

	tests := []struct {
		input inputStruct
		want  []byte
	}{
		{
			input: inputStruct{
				a: []byte{0b0001_1100, 0b0000_0001, 0b0001_0001, 0b0000_0000, 0b0001_1111, 0b0000_0001, 0b0000_0001, 0b0000_0000, 0b0000_0110, 0b0001_1010, 0b0000_0010, 0b0100_1011, 0b0101_0011, 0b0101_0011, 0b0101_0000, 0b0000_1001, 0b0001_1000, 0b0001_1100},
				b: []byte{0b0110_1000, 0b0110_1001, 0b0111_0100, 0b0010_0000, 0b0111_0100, 0b0110_1000, 0b0110_0101, 0b0010_0000, 0b0110_0010, 0b0111_0101, 0b0110_1100, 0b0110_1100, 0b0010_0111, 0b0111_0011, 0b0010_0000, 0b0110_0101, 0b0111_1001, 0b0110_0101},
			},
			want: []byte{0b0111_0100, 0b0110_1000, 0b0110_0101, 0b0010_0000, 0b0110_1011, 0b0110_1001, 0b0110_0100, 0b0010_0000, 0b0110_0100, 0b0110_1111, 0b0110_1110, 0b0010_0111, 0b0111_0100, 0b0010_0000, 0b0111_0000, 0b0110_1100, 0b0110_0001, 0b0111_1001},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%#v", tt.input), func(t *testing.T) {
			got, err := XOR(tt.input.a, tt.input.b)
			if !bytes.Equal(got, tt.want) || err != nil {
				t.Errorf("got: '%v', want: '%v', err: %v", got, tt.want, err)
			}
		})
	}
}

func TestToHex(t *testing.T) {
	tests := []struct {
		input []byte
		want  string
	}{
		{[]byte{0b0111_0100, 0b0110_1000, 0b0110_0101, 0b0010_0000, 0b0110_1011, 0b0110_1001, 0b0110_0100, 0b0010_0000, 0b0110_0100, 0b0110_1111, 0b0110_1110, 0b0010_0111, 0b0111_0100, 0b0010_0000, 0b0111_0000, 0b0110_1100, 0b0110_0001, 0b0111_1001}, "746865206B696420646F6E277420706C6179"},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%#v", tt.input), func(t *testing.T) {
			got, err := HexToString(tt.input)
			if got != tt.want || err != nil {
				t.Errorf("got: '%v', want: '%v', err: %v", got, tt.want, err)
			}
		})
	}
}

func TestDistance(t *testing.T) {

	type testStruct struct {
		a string
		b string
	}
	tests := []struct {
		input testStruct
		want  int
	}{
		{testStruct{a: "this is a test", b: "wokka wokka!!!"}, 37},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%#v", tt.input), func(t *testing.T) {
			got, err := Distance(tt.input.a, tt.input.b)
			if got != tt.want || err != nil {
				t.Errorf("got: '%v', want: '%v', err: %v", got, tt.want, err)
			}
		})
	}
}
