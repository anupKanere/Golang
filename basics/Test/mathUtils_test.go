package main

import "testing"

func TestAdd(t *testing.T) {
	res := add(2, 4)
	expected := 6
	if res != expected {
		t.Errorf("Expected %d got %d ", expected, res)
	}
}
