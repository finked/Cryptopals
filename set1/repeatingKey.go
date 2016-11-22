package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	// filename := "./data/6.txt"
	// keysize := getKeysize(filename)
	// fmt.Println(keysize)

	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	letter := "58"
	data, _ := hex.DecodeString(input)
	data2, _ := hex.DecodeString(letter)
	result := XorLetter(data, data2)
	sol := string(result[:])
	fmt.Println(sol)
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

func hammingDist(input1, input2 []byte) int {
	/*
	 *	Calculate hamming distance between two byte arrays
	 *	(Count each different bit)
	 */
	count := 0
	for idx, b1 := range input1 {
		b2 := input2[idx]
		xor := b1 ^ b2

		for x := xor; x > 0; x >>= 1 {
			if int(x&1) == 1 {
				count++
			}
		}
	}
	return count
}

func getTwoParts(length int, filename string) ([]byte, []byte) {
	f, _ := os.Open(filename)
	r := bufio.NewReader(f)

	data, _, _ := r.ReadLine()
	// fmt.Println(len(data))
	return data[0:length], data[length : 2*length]
}

func getKeysize(filename string) int {
	min := 40
	imin := 0
	for i := 2; i < 40; i++ {
		data1, data2 := getTwoParts(i, filename)
		val := hammingDist(data1, data2) / i
		if min > val {
			min = val
			imin = i
		}
	}
	return imin
}
