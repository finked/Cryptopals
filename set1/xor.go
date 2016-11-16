package main

import "encoding/hex"

// import "fmt"

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

// func XorByte(input []byte, input2 []byte) string {
// data, _ := hex.DecodeString(input)
// data2, _ := hex.DecodeString(input2)

// res := make([]byte, len(data))
// for i := 0; i < len(data) && i < len(data2); i++ {
// res[i] = data[i] ^ data2[i]
// }

// sol := hex.EncodeToString(res)
// return sol
// }

func XorLetter(data []byte, letter []byte) []byte {
	res := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(letter); j++ {
			res[i] = data[i] ^ letter[j]
		}
	}

	return res
}

// func XorLetter(input string, letter string) []byte {
// data, _ := hex.DecodeString(input)
// data2, _ := hex.DecodeString(letter)

// res := make([]byte, len(data))
// for i := 0; i < len(data); i++ {
// for j := 0; j < len(data2); j++ {
// res[i] = data[i] ^ data2[j]
// }
// }

// return res
// }
