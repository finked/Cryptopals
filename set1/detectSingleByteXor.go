package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	// "reflect"
)

func readFile(filename string) []byte {
	f, _ := os.Open(filename)
	r := bufio.NewReader(f)

	line, isPrefix, err := r.ReadLine()
	bestScore := 500.0
	var bestLine []byte
	ind := 0
	bestInd := 0
	// TODO(DF): Add index idx in loop and remove ind
	for err == nil && !isPrefix {
		buf, _ := hex.DecodeString(string(line[:]))
		// fmt.Println(string(buf[:]))
		sol, score := getBestSol(buf)
		// fmt.Println(score)
		// letter, _ := hex.DecodeString(string(sol))
		letter, _ := hex.DecodeString(fmt.Sprintf("%x", string(sol)))
		// fmt.Println(letter)
		res := XorLetter(buf, letter)
		fmt.Printf("%d:%f,%s\n", ind, score, string(res[:]))

		if score < bestScore {
			bestScore = score
			bestLine = res
			bestInd = ind
			// fmt.Println(ind)
		}

		line, isPrefix, err = r.ReadLine()
		ind++
	}

	// sol := hex.EncodeToString(bestLine)
	fmt.Printf("Best line is %v with score %v\n", string(bestLine[:]), bestScore)
	fmt.Printf("Best index %d\n", bestInd)

	return bestLine[:]
}
