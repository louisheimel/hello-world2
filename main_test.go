package main

import "testing"

func TestSubtract(t *testing.T) {
	res := Subtract(2)
	if res != 4 {
		t.Error()
	}
}