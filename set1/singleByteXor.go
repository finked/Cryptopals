package main

import "fmt"

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	for i := 0; i < 127; i++ {
		letter := fmt.Sprintf("%x", i)
		res := XorLetter(input, letter)
		fmt.Printf("i = %v, res = %s\n", i, string(res[:]))
	}
}

func scoreEngText(input []byte) {
	// Count same letters and compare to most used english letters
	for i := 0; i < len(input); i++ {
	}

	// Return estimated error
}

func singleByteXor(input string) string {
	// Most used letters ETAOIN SHRDLU
	// min := 140
	res := "Blub"
	for i := 0; i < 127; i++ {
		// letter := fmt.Sprintf("%x", i)
		// res = XorLetter(input, letter)

		// Calculate propability
	}

	return string(res[:])
}

func transformToAscii(input string) string {
	return "buf"
}
