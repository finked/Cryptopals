package main

import (
	"encoding/hex"
	// "fmt"
)

func Xor(input string, input2 string) string {
	data, _ := hex.DecodeString(input)
	data2, _ := hex.DecodeString(input2)

	res := make([]byte, len(data))
	for i := 0; i < len(data) && i < len(data2); i++ {
		res[i] = data[i] ^ data2[i]
	}

	sol := hex.EncodeToString(res)
	return sol
}

// func XorLetter(data, letter []byte) []byte {
// res := make([]byte, len(data))
// letlen := len(letter)
// if letlen == 0 {
// return []byte{0}
// }
// for i := 0; i < len(data); i++ {
// res[i] = data[i] ^ letter[i%letlen]
// fmt.Println(res[i])
// }
// return res
// }

func XorLetter(data []byte, letter []byte) []byte {
	// fmt.Println(len(data))
	// fmt.Println(len(letter))
	res := make([]byte, len(data))
	letlen := len(letter)
	if letlen == 0 {
		letlen = 1
	}
	for i := 0; i < len(data)/letlen+1; i++ {
		for j := 0; j < len(letter); j++ {
			if i*letlen+j == len(data) {
				return res
			}
			res[i*letlen+j] = data[i*letlen+j] ^ letter[j]
		}
	}

	return res
}
