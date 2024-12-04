package main

import (
	"cryptopals-go/base64"
	"log"
	"os"
	"strings"
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
	setOneChallengeFour()
	setOneChallengeFive()
}

func setOneChallengeOne() {

	heading := "Set One / Challenge One failed"

	hexBytes, err := HexToBytes("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	encoding := base64.Encode(string(hexBytes))

	if encoding != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
		log.Fatalf("%s: %v", heading, "incorrect value")
	}

	log.Printf("challenge 1 | answer: %v", encoding)
}

func setOneChallengeTwo() {

	heading := "Set One / Challenge Two failed"

	bytesOne, err := HexToBytes("1c0111001f010100061a024b53535009181c")
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	bytesTwo, err := HexToBytes("686974207468652062756c6c277320657965")
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	operation, err := XOR(bytesOne, bytesTwo)
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	hex, err := HexToString(operation)
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	if hex != "746865206B696420646F6E277420706C6179" {
		log.Fatalf("%s: %v", heading, "incorrect value")
	}

	log.Printf("challenge 2 | answer: %v", hex)
}

func setOneChallengeThree() {

	heading := "Set One / Challenge Three failed"

	bytesValue, err := HexToBytes("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
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

	// Answer: Cooking MC's like a pound of bacon
	for i, value := range Score(decryptValues)[0:5] {
		log.Printf("challenge 3 | number %d: %v\n", i+1, value)
	}
}

func setOneChallengeFour() {

	heading := "Set One / Challenge Four failed"

	contents, err := os.ReadFile("testdata/set-one-challenge-four.txt")
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	lines := strings.Split(string(contents), "\n")
	if len(lines) <= 0 {
		log.Fatalf("%s: %v", heading, "no rows in file")
	}

	decBytes, err := HexToBytes(lines[0])
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	// Score for each iteration of decryption
	iterScores := make([]string, 256)
	// Top scores for each iteration
	topScores := make([]string, 0, 5*len(lines))

	var length int
	var keyExpand []byte

	for _, encValue := range lines {

		decBytes, err = HexToBytes(encValue)
		if err != nil {
			log.Fatalf("%s: %v", heading, err)
		}

		length = len(decBytes)
		keyExpand = make([]byte, length)

		for i := range 256 {

			for j := range length {
				keyExpand[j] = byte(i)
			}

			decryptBytes, err := XOR(decBytes, keyExpand)
			if err != nil {
				log.Fatalf("%s: %v", heading, err)
			}

			iterScores[i] = ToASCII(decryptBytes)
		}

		topScores = append(topScores, Score(iterScores)[0:5]...)
	}

	// Answer: Now that the party is jumping
	for i, value := range Score(topScores)[0:5] {
		log.Printf("challenge 4 | number %d: %v\n", i+1, value)
	}
}

func setOneChallengeFive() {

	heading := "Set One / Challenge Five failed"

	message := `Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`

	key := []rune{'I', 'C', 'E'}

	length := len(message)
	keyExpand := make([]byte, length)

	for i := range len(message) {
		keyExpand[i] = byte(key[i%3])
	}

	encBytes, err := XOR([]byte(message), keyExpand)
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	hex, err := HexToString(encBytes)
	if err != nil {
		log.Fatalf("%s: %v", heading, err)
	}

	if hex != strings.ToUpper("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f") {
		log.Fatalf("%s: %v", heading, "incorrect value")
	}

	log.Printf("challenge 5 | answer: %v", hex)
}
