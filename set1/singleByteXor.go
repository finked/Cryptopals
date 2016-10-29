package main

import "fmt"
import "math"
import "strconv"

// import "encoding/binary"

func getBestSol(input string) (string, float64) {
	var bestSol int
	bestScore := float64(len(input))
	for i := 65; i < 122; i++ {
		num := fmt.Sprintf("%x", i)
		// fmt.Println(num)
		res := XorLetter(input, num)
		// fmt.Printf("Solution:\nres = %s\n", string(res[:]))
		// fmt.Printf("byte = %v\n", res)
		score := scoreEngText(res)
		// fmt.Printf("Score is %v with letter %v\n\n", score, string(i))

		if score < bestScore {
			bestScore = score
			bestSol = i
		}
	}

	// fmt.Printf("Score is %v\n", bestScore)
	// fmt.Printf("Best solution is with letter %v\n", string(bestSol))
	// res := XorLetter(input, fmt.Sprintf("%x", bestSol))
	// fmt.Printf("Solution:\nres = %s\n", string(res[:]))

	return string(bestSol), bestScore
}

func scoreEngText(input []byte) float64 {
	// Count same letters
	var ascii [128]int
	// fmt.Printf("len = %d cap = %d %v\n", len(ascii), cap(ascii), ascii)

	for i := 0; i < len(input); i++ {
		num := fmt.Sprintf("%d", input[i])
		inum, _ := strconv.Atoi(num)
		// fmt.Printf("num = %v, inum = %v\n", num, inum)
		for j := 0; j < 128; j++ {
			// fmt.Println(num)
			if j == inum {
				ascii[j]++
				// fmt.Printf("We got it. New count is %v for %v\n", ascii[j], j)
			}
		}
	}

	// Calculate error
	// TODO(DF): 	Remove whitespace number from N. Add upper case letters for start
	// 				Remove constant term at the end
	N := len(input)
	errval := [26]float64{8.167, 1.492, 2.782, 4.253, 12.702, 2.228, 2.015, 6.094, 6.966,
		0.153, 0.772, 4.025, 2.406, 6.749, 7.507, 1.929, 0.095, 5.987, 6.327, 9.056,
		2.758, 0.978, 2.360, 0.150, 1.974, 0.074}

	var err [26]float64
	for i := 0; i < 26; i++ {
		err[i] = math.Abs(errval[i]*float64(N)/100 - (float64(ascii[97+i])))
		// err[i] = math.Abs(errval[i]*float64(N)/100 - (float64(ascii[65+i] + ascii[97+i])))
	}

	calcerr := sliceSum(err)

	// Return estimated error
	return calcerr
}

func sliceSum(slice [26]float64) float64 {
	var sum float64
	for i := 0; i < len(slice); i++ {
		sum += slice[i]
	}
	return sum
}
