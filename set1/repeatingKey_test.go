package main

import (
	"fmt"
	"testing"
)

func TestRepeatingKey(t *testing.T) {
	input := "Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"
	data := []byte(input)
	fmt.Println(data)

	res := string(repeatingKey(data)[:])

	// Solution:
	sol := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"

	if res != sol {
		t.Errorf("Expected: \n%s got \n%s", sol, res)
	}
}

func TestHammingDist(t *testing.T) {
	input1 := "this is a test"
	input2 := "wokka wokka!!!"

	res := hammingDist([]byte(input1), []byte(input2))
	if res != 37 {
		t.Errorf("Expected: 37 got %d", res)
	}
}

func TestGetTwoPart(t *testing.T) {
	filename := "./data/6.txt"
	res1, res2 := getTwoParts(2, filename)
	if string(res1) != "HU" && string(res2) != "If" {
		t.Errorf("Expected: HU and If got %s, %s", string(res1), string(res2))
	}
}

func TestGetKeysize(t *testing.T) {
	// input := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272"
	// TODO(DF): Change functions to handle input byte arrays and split file reading to other function
}
