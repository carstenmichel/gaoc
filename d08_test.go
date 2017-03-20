package main

import "testing"

// TestMain is test for d1
func TestD08(t *testing.T) {
	leds := d08()
	if leds != 128 {
		t.Fatalf("Result %v is not ok \n", leds)
	}

}
