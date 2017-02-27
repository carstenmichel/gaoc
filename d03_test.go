package main

import "testing"

// TestMain is test for d1
func TestD03(t *testing.T) {
	first := d03()
	if first != 993 {
		t.Fatalf("First Puzzle wrong with %v\n", first)
	}
	second := d03Part2()
	if second != 1849 {
		t.Fatalf("Second Puzzle wrong with %v\n", first)
	}
}
