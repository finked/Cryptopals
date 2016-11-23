package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	// "strconv"
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
	f, _ := os.Open(filename)
	r := bufio.NewReader(f)

	data, _, _ := r.ReadLine()
	res := getParts(2, 2, data)
	if string(res[0]) != "HU" && string(res[1]) != "If" {
		t.Errorf("Expected: HU and If got %s, %s", string(res[0]), string(res[1]))
	}
}

func TestGetKeysize(t *testing.T) {
	input := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	data, _ := hex.DecodeString(input)
	keysize := getKeysize(3, data)

	if keysize != 3 {
		t.Errorf("Expected: 3 got %d", keysize)
	}
}

func TestSplitData(t *testing.T) {
	input := []byte{'a', 'b', 'c', 'd'}

	res := splitData(2, input)

	if res[0][0] != 'a' || res[0][1] != 'b' || res[1][0] != 'c' || res[1][1] != 'd' {
		t.Errorf("Expected: ab cd")
	}
}

func TestTransposeData(t *testing.T) {
	input := make([][]byte, 2)
	input[0] = []byte{'a', 'b'}
	input[1] = []byte{'c', 'd'}

	res := transposeData(input)

	if res[0][0] != 'a' || res[0][1] != 'c' || res[1][0] != 'b' || res[1][1] != 'd' {
		t.Errorf("Expected: ac bd")
	}
}

func TestKeyCalc(t *testing.T) {
	input := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	data, _ := hex.DecodeString(input)

	// Split and transpose data
	split := splitData(3, []byte(data))
	tsplit := transposeData(split)

	// Get key
	key := make([]byte, 3)
	for i := 0; i < 3; i++ {
		bestSol, _ := getBestSol(tsplit[i])
		// ikey := bestSol
		key[i] = byte(bestSol)
	}

	if string(key) != "ICE" {
		t.Errorf("Expected: ICE, got %s", string(key))
	}
}
