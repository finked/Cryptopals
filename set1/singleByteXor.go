package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"sort"
	"strconv"
)

func getBestSol(data []byte) (int, float64) {
	var bestSol int
	bestScore := 0.0
	for i := 65; i < 122; i++ {
		num := fmt.Sprintf("%x", i)
		// data, _ := hex.DecodeString(input)
		data2, _ := hex.DecodeString(num)
		res := XorLetter(data, data2)
		score := scoreEngText(res)

		if score > bestScore {
			bestScore = score
			bestSol = i
		}
	}

	// res := XorLetter(input, fmt.Sprintf("%x", bestSol))
	// fmt.Printf("Solution:\nres = %s\n", string(res[:]))

	return bestSol, bestScore
}

func scoreEngText(input []byte) float64 {
	// Count same letters
	var ascii [128]int
	// fmt.Printf("len = %d cap = %d %v\n", len(ascii), cap(ascii), ascii)

	// fmt.Println(input)

	for i := 0; i < len(input); i++ {
		// fmt.Println(input[i])
		num := fmt.Sprintf("%d", input[i])
		// fmt.Println(num)
		inum, _ := strconv.Atoi(num)
		// fmt.Println(inum)
		ascii[inum]++
	}

	// Calculate error
	N := len(input) - ascii[32]
	// N := len(input)
	// N := float64(sliceIntSum(ascii[65:122], 26))
	errval := [26]float64{8.167, 1.492, 2.782, 4.253, 12.702, 2.228, 2.015, 6.094, 6.966,
		0.153, 0.772, 4.025, 2.406, 6.749, 7.507, 1.929, 0.095, 5.987, 6.327, 9.056,
		2.758, 0.978, 2.360, 0.150, 1.974, 0.074}

	var err [26]float64
	for i := 0; i < 26; i++ {
		err[i] = math.Abs((errval[i] * float64(ascii[97+i])) / float64(N))
		// err[i] = math.Sqrt(math.Abs(errval[i]*float64(N)/100 -
		// (float64(ascii[97+i]))))
		// err[i] = math.Abs(errval[i]*float64(N)/100 - (float64(ascii[65+i] + ascii[97+i])))
	}

	calcerr := sliceSum(err)

	// Return estimated error
	return calcerr
}

func countFrequencys(input []byte) []int {
	var ascii []int
	ascii = make([]int, 26, 26)
	for i := 0; i < len(input); i++ {
		num := fmt.Sprintf("%d", input[i])
		inum, _ := strconv.Atoi(num)
		if inum > 96 && inum < 123 {
			ascii[inum-96]++
		}
	}

	return ascii
}

func evalLetFreq(letfreq []int) string {
	sort.Ints(letfreq)

	return "X"
}

func sliceSum(slice [26]float64) float64 {
	var sum float64
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	return sum
}

func sliceIntSum(slice []int) int {
	var sum int
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	return sum
}
