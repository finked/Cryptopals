package main

import "fmt"
import "encoding/hex"

func main() {
	input := "1c0111001f010100061a024b53535009181c"
	// solution 746865206b696420646f6e277420706c6179
	data, _ := hex.DecodeString(input)

	input2 := "686974207468652062756c6c277320657965"

	data2, _ := hex.DecodeString(input2)

	//var res []byte
	res := make([]byte, len(data))
	//res = res[len(data)]
	for i := 0; i < len(data); i++ {
		res[i] = data[i] ^ data2[i]
	}

	sol := hex.EncodeToString(res)
	fmt.Println(sol)

}
