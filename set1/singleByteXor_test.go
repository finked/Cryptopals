package main

import "testing"

func TestSingleByteXor(t *testing.T) {
	data := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	result := singleByteXor(data)
	// X (int 88) is the correct letter
	if result != "X" {
		t.Errorf("Expected .X, got $s", result)
	}
}
