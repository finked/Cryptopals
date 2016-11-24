package main

import (
	"testing"
)

func TestExpandKey(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")
	res := expandKey(key)

	if string(res) != "TEST" {
		t.Errorf("Expected: ..., got %s", string(res))
	}
}
