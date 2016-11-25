package main

import (
	"testing"
)

func TestExpandKey(t *testing.T) {
	key := []byte("YELLOW SUBMARINE")
	res := expandKey(key, 1)

	if string(res) != "YELSOW AUBMERINL" {
		t.Errorf("Expected: YELSOW AUBMERINL got %s", string(res))
	}

	res2 := expandKey(key, 3)
	if string(res2) != "YELEOW LUBMSRINA" {
		t.Errorf("Expected: YELEOW LUBMSRINA got %s", string(res))
	}
}
