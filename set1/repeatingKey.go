package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	filename := "./data/6.txt"
	keysize := getKeysize(filename)
	fmt.Println(keysize)
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
	min := 40.0
	imin := 0
	for i := 2; i < 41; i++ {
		data1, data2 := getTwoParts(i, filename)
		val1 := hammingDist(data1, data2)
		val := float64(val1) / float64(i)
		// fmt.Printf("ham = %d, val = %f, ind = %d\n", val1, val, i)
		if min > val {
			min = val
			imin = i
		}
	}
	return imin
}
