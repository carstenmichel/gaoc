package main

import "testing"

// TestMain is test for d1
func TestD04(t *testing.T) {
	first, sector := d04()
	if first != 409147 {
		t.Fatalf("First Puzzle wrong with %v\n", first)
	}
	if sector != 991 {
		t.Fatalf("Secret sector wrong with %v\n", sector)
	}

}
