package main

import "fmt"
import "strconv"

// import "encoding/binary"

func main() {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"

	res := XorLetter(input, "58")
	fmt.Printf("Solution:\nlen(res) = %v, res = %s\n", len(res), string(res[:]))
	scoreEngText(res)
	// for i := 0; i < 128; i++ {
	// letter := fmt.Sprintf("%x", i)
	// res := XorLetter(input, letter)
	// fmt.Printf("i = %v, res = %s\n", i, string(res[:]))
	// }

	// buf := []byte{54, 54, 49, 49}
	fmt.Printf("res = %v\n", res)

	x, _ := strconv.Atoi(string(res[5:6]))
	fmt.Println(x)
}

func scoreEngText(input []byte) {
	// Count same letters and compare to most used english letters
	// var ascii [127]int
	// fmt.Printf("len = %d cap = %d %v\n", len(ascii), cap(ascii), ascii)

	// bs := make([]byte, 4)

	// var j uint32
	// fmt.Println(len(input))
	for i := 0; i < len(input); i++ {
		// fmt.Printf("letter = %v\n", input[i])
		for j := 0; j < 128; j++ {
			// fmt.Sprintf("%v\n", input[i])
			// test := fmt.Sprintln("%v", input[i])
			// if j == test {
			// fmt.Println("We got it\n")
			// }
			// test := binary.BigEndian.Uint64(input[i : i+1])
			// fmt.Println(test)
			// if input == bs {
			// ascii[j]++
			// }
		}
	}
	// Print found values
	// fmt.Printf("e = %x\n", ascii[102])
	// fmt.Printf("t = %x\n", ascii[116])
	// fmt.Printf("a = %x\n", ascii[97])

	// Return estimated error
}
