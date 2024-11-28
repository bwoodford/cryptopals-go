package main

import (
	"cryptopals-go/base64"
	"log"
)

func main() {
	setOneChallengeOne()
	setOneChallengeTwo()
}

func setOneChallengeOne() {

	heading := "Set One / Challenge One failed"

	encoding := base64.Encode("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

	if encoding != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		log.Fatalf("%s: %v", heading, "incorrect value")
	}
}

func setOneChallengeTwo() {

	heading := "Set One / Challenge Two failed"

	bytesOne, err := ToBytes("1c0111001f010100061a024b53535009181c")
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	bytesTwo, err := ToBytes("686974207468652062756c6c277320657965")
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	operation, err := XOR(bytesOne, bytesTwo)
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	hex, err := ToHex(operation)
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	if hex != "746865206B696420646F6E277420706C6179" {
		log.Fatalf("%s: %v", heading, "incorrect value")
	}
}
