package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	// "reflect"
)

func main() {
	readFile("./data/4.txt")

	// Solution:
	// data, _ := hex.DecodeString("7b5a4215415d544115415d5015455447414c155c46155f4058455c5b523f")
	// val, _ := hex.DecodeString("35")
	// res := XorLetter(data, val)
	// fmt.Println(string(res[:]))
}

func readFile(filename string) []byte {
	f, _ := os.Open(filename)
	r := bufio.NewReader(f)

	line, isPrefix, err := r.ReadLine()
	fmt.Println("Read Line:")
	fmt.Println(line)
	// fmt.Println(reflect.TypeOf(line))
	bestScore := float64(len(line))
	var bestLine []byte
	for err == nil && !isPrefix {
		// stest := fmt.Sprintf("%s", line[:])
		// fmt.Println(stest)
		// s := teststring
		// s := string(line[:])
		sol, score := getBestSol(line)
		// fmt.Println(sol)
		// letter, _ := hex.DecodeString(string(sol))
		letter, _ := hex.DecodeString(fmt.Sprintf("%x", string(sol)))
		fmt.Println(letter)
		res := XorLetter(line, letter)
		fmt.Printf("res = %s\n", string(res[:]))

		if bestScore > score {
			bestScore = score
			bestLine = line
		}

		line, isPrefix, err = r.ReadLine()
	}

	// sol := hex.EncodeToString(bestLine)
	fmt.Printf("Best line is %v with score %v", string(bestLine[:]), bestScore)
	return bestLine[:]
}

// func getBestByteSol(input []byte) (string, float64) {
// var bestSol int
// bestScore := 0.0
// for i := 65; i < 122; i++ {
// num := fmt.Sprintf("%x", i)
// res := XorLetter(input, num)
// score := scoreEngText(res)
// // fmt.Printf("Score is %v with letter %v\n\n", score, string(i))

// if score > bestScore {
// bestScore = score
// bestSol = i
// }
// }

// // res := XorLetter(input, fmt.Sprintf("%x", bestSol))
// // fmt.Printf("Solution:\nres = %s\n", string(res[:]))

// return string(bestSol), bestScore
// }

// func ByteXor(input []byte, input2 string) string {
// data, _ := hex.DecodeString(input)
// data2, _ := hex.DecodeString(input2)

// res := make([]byte, len(data))
// for i := 0; i < len(data) && i < len(data2); i++ {
// res[i] = data[i] ^ data2[i]
// }

// sol := hex.EncodeToString(res)
// return sol
// }
