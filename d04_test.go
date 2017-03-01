package main

import "testing"

// TestMain is test for d1
func TestD04(t *testing.T) {
	first := d04()
	if first != 409147 {
		t.Fatalf("First Puzzle wrong with %v\n", first)
	}

}
