package main

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	res := readFile("./data/4.txt")
	if string(res) != "Now that the party is jumping" {
		t.Errorf("Expected: Now that the party is jumping, got $s", string(res))
	}
}
