package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Read file
	filename := "./data/6.txt"
	f, _ := os.Open(filename)
	r := bufio.NewReader(f)
	data, _, _ := r.ReadLine()

	// Encode it with base64
	decData, _ := base64.StdEncoding.DecodeString(string(data))

	// Get most likely keysize
	keysize := getKeysize(decData)
	fmt.Println(keysize)

	// Read in complete file
	test, _ := ioutil.ReadFile(filename)
	blocks := splitData(keysize, test)

	// Transpose the data
	tblock := transposeData(blocks)

	// Solve each block for one letter
	key := make([]byte, keysize)
	for i := 0; i < keysize; i++ {
		bestSol, _ := getBestSol(tblock[i])
		key[i] = byte(fmt.Sprintf("%x", bestSol))
	}

	// Use key to decrypt the file
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

func getTwoParts(length int, data []byte) ([]byte, []byte) {
	// fmt.Println(len(data))
	if 2*length <= len(data) {
		return data[0:length], data[length : 2*length]
	} else {
		return nil, nil
	}
}

func getKeysize(data []byte) int {
	min := float64(len(data))
	imin := 0

	for i := 2; i < 41; i++ {
		if len(data) < i*2 {
			return imin
		}
		data1, data2 := getTwoParts(i, data)
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

func splitData(length int, data []byte) [][]byte {
	/*
	 *	Split data in blocks with given length
	 */
	dlen := len(data)/length + 1
	res := make([][]byte, dlen)
	for i, _ := range dlen {
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
	for i, _ := range dlen {
		// res[i] := make([]byte, len(data))
		for j, _ := range len(data) {
			res[i][j] = data[j][i]
		}
	}
	return res
}
