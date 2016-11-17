package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	first := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"

	fmt.Println(len(first))
	res := repeatingKey([]byte(first))
	fmt.Println(len(res))
	fmt.Println(res)
}

func repeatingKey(input []byte) string {
	// Encrypt with ICE
	letter := []byte{'I', 'C', 'E'}
	res := XorLetter(input, letter)
	fmt.Println(len(res))
	res2 := hex.EncodeToString(res)

	// Return data
	return res2
}
