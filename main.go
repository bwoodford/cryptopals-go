package main

import (
	"cryptopals-go/base64"
	"log"
	"os"
)

func main() {
	// Open or create the log file
	file, err := os.OpenFile("pals.log", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Set the log output to the file
	log.SetOutput(file)

	setOneChallengeOne()
	setOneChallengeTwo()
	setOneChallengeThree()
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

func setOneChallengeThree() {

	heading := "Set One / Challenge Three failed"

	bytesValue, err := ToBytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	length := len(bytesValue)
	keyExpand := make([]byte, length)
	decryptValues := make([]string, 256)

	for i := range 256 {

		for j := range length {
			keyExpand[j] = byte(i)
		}

		decryptBytes, err := XOR(bytesValue, keyExpand)
		if err != nil {
			log.Fatalf("%s: %v", heading, err)
		}

		decryptValues[i] = ToASCII(decryptBytes)
	}

	for i, value := range Score(decryptValues)[0:5] {
		log.Printf("number %d: %v\n", i+1, value)
	}

	/*
		secretKey := byte('X')

		for j := range length {
			keyExpand[j] = secretKey
		}

		decryptBytes, err := XOR(bytesValue, keyExpand)
		if err != nil {
			log.Fatalf("%s: %v", heading, err)
		}

		if ToASCII(decryptBytes) != "Cooking MC's like a pound of bacon" {
			log.Fatalf("%s: %v", heading, "incorrect value")
		}
	*/
}

func setOneChallengeFour() {

}
