package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math"
)

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

func getParts(length int, number int, data []byte) [][]byte {
	if 2*length <= len(data) {
		ret := make([][]byte, number)
		for i := 0; i < number; i++ {
			ret[i] = data[length*i : length*(i+1)]
		}
		return ret
	} else {
		return nil
	}
}

// TODO(DF): Return k best key values instead of the best one
// TODO(DF): Use each block with each block to compare
func getKeysize(num int, data []byte) int {
	min := float64(len(data))
	imin := 0

	for i := 2; i < 41; i++ {
		if len(data) < i*num*2 {
			fmt.Println("Abort because of short data")
			return imin
		}
		data := getParts(i, num*2, data)
		dist := 0
		for j := 0; j < num; j++ {
			dist += hammingDist(data[j*2], data[j*2+1])
		}
		val := float64(dist) / float64(num*i)
		// fmt.Printf("ham = %d, val = %f, ind = %d\n", dist, val, i)
		if min > val {
			min = val
			imin = i
		}
	}
	return imin
}

func splitData(length int, data []byte) [][]byte {
	/*
	 *	Split data in blocks with given length
	 */
	dlen := int(math.Ceil(float64(len(data)) / float64(length)))
	// TODO(DF): Have to handle rest if it doesn't fit
	res := make([][]byte, dlen)
	for i, _ := range res {
		end := (i + 1) * length
		if end > len(data) {
			end = len(data)
		}
		res[i] = data[i*length : end]
	}
	return res
}

func transposeData(data [][]byte) [][]byte {
	/*
	 *	Transpose given data set
	 */
	dlen := len(data[0])
	res := make([][]byte, dlen)
	for i := 0; i < dlen; i++ {
		// fmt.Println(i)
		res[i] = make([]byte, len(data))
		// fmt.Println(len(res[i]))
		for j := 0; j < len(data); j++ {
			// fmt.Printf("j = %d, i = %d\n", j, i)
			if len(data[j]) > i {
				res[i][j] = data[j][i]
			}
		}
	}
	return res
}

func breakRepeatingKey(filename string) string {
	// Read file
	data, _ := ioutil.ReadFile(filename)

	// Decode it with base64
	decData, _ := base64.StdEncoding.DecodeString(string(data))

	// Get most likely keysize
	keysize := getKeysize(8, decData)
	// fmt.Println(keysize)

	// Read in complete file
	blocks := splitData(keysize, decData)

	// Transpose the data
	tblock := transposeData(blocks)

	// Solve each block for one letter
	key := make([]byte, keysize)
	for i := 0; i < keysize; i++ {
		// fmt.Printf("index %d, text = %s\n", i, string(tblock[i][0:40]))
		bestSol, _ := getBestSol(tblock[i])
		// fmt.Printf("bestSol = %d, %s\n", bestSol, string(bestSol))
		key[i] = byte(bestSol)
	}
	// fmt.Printf("i = %d, key = %s\n", i, string(key))
	// fmt.Printf("key = %s\n", string(key))

	// Use key to decrypt the file
	res := XorLetter(decData, key)
	fmt.Println(string(res))

	return string(key)
}
