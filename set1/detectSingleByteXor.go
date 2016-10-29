package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile("4.txt")
}

func readFile(filename string) {
	f, _ := os.Open(filename)
	r := bufio.NewReader(f)

	line, isPrefix, err := r.ReadLine()
	bestScore := float64(len(line))
	var bestLine string
	for err == nil && !isPrefix {
		s := string(line)
		sol, score := getBestSol(s)
		res := XorLetter(s, fmt.Sprintf("%x", sol))
		fmt.Printf("res = %s\n", string(res[:]))

		if bestScore > score {
			bestScore = score
			bestLine = s
		}

		line, isPrefix, err = r.ReadLine()
	}

	fmt.Printf("Best line is %v with score %v", bestLine, bestScore)
}
