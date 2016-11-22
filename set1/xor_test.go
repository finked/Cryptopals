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
	letter := "58" // hex for X
	data, _ := hex.DecodeString(input)
	data2, _ := hex.DecodeString(letter)
	result := XorLetter(data, data2)
	sol := string(result[:])
	if sol != "Cooking MC's like a pound of bacon" {
		t.Errorf("Expected .Cooking MC's like a pound of bacon, got $s", sol)
	}
}

func TestXorLetterLong(t *testing.T) {
	input := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	letter := "494345" // hex for ICE
	data, _ := hex.DecodeString(input)
	data2, _ := hex.DecodeString(letter)
	result := XorLetter(data, data2)
	if string(result) != "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal" {
		t.Errorf("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal, got $s", string(result))
	}
}
