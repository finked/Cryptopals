package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Read file
	filename := "./data/7.txt"
	data, _ := ioutil.ReadFile(filename)

	fmt.Println(data)
}
