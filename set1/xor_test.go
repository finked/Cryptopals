package main

import "testing"

func TestXor(t *testing.T) {
	input := "1c0111001f010100061a024b53535009181c"
	input2 := "686974207468652062756c6c277320657965"
	result := Xor(input, input2)
	if result != "746865206b696420646f6e277420706c6179" {
		t.Errorf("Expected: 746865206b696420646f6e277420706c6179", result)
	}
}
