package main

import (
	"encoding/hex"
	"testing"
)

func TestXor(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	result := Xor(input, input2)
	if result != "746865206b696420646f6e277420706c6179" {
		t.Errorf("Expected: 746865206b696420646f6e277420706c6179", result)
	}
}

func TestXorLetter(t *testing.T) {
	input := "1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"
	letter := "58"
	data, _ := hex.DecodeString(input)
	data2, _ := hex.DecodeString(letter)
	result := XorLetter(data, data2)
	sol := string(result[:])
	if sol != "Cooking MC's like a pound of bacon" {
		t.Errorf("Expected .Cooking MC's like a pound of bacon, got $s", sol)
	}
}
