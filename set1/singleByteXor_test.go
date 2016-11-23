package main

import (
	"encoding/hex"
	"testing"
)

func TestSingleByteXor(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	data, _ := hex.DecodeString(input)
	result, _ := getBestSol(data)
	if string(result) != "X" {
		t.Errorf("Expected .X, got $s", result)
	}
}

func TestScoreEngText(t *testing.T) {
	input := "     "
	data, _ := hex.DecodeString(input)
	res := scoreEngText(data)

	if res != 0 {
		t.Errorf("Expected 0, got %d", res)
	}
}
