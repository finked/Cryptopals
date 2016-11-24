package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Read file
	filename := "./data/7.txt"
	data, _ := ioutil.ReadFile(filename)

	// key := "YELLOW SUBMARINE"

	fmt.Println(data)

	// Create roundkeys

	// Xor

	// ShiftRows

}

func expandKey(key []byte) []byte {
	/*
	 *	rotate 4th column (4x4 key) 1 step up
	 */
	return nil
}

func invSubBox(key []byte) []byte {
	return nil
}

func generateRoundKeys(key []byte) [][]byte {
	return nil
}

func shiftBlocks(data [][]byte) {
}
