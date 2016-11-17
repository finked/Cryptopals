package main

import (
	"encoding/hex"
	"testing"
)

func TestReadFile(t *testing.T) {
	res := readFile("./data/4.txt")
	if string(res) != "Now that the party is jumping\n" {
		t.Errorf("Expected: Now that the party is jumping\n got %s", string(res))
	}
}

func TestReadCorrectLine(t *testing.T) {
	data, _ := hex.DecodeString("7b5a4215415d544115415d5015455447414c155c46155f4058455c5b523f")
	res, _ := getBestSol(data)
	// fmt.Println(string(res[:]))
	if string(res) != "5" {
		t.Errorf("Expected: 5, got %s", string(res))
	}
}
