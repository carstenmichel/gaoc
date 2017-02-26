package main

import "testing"

// TestMain is test for d1
func TestD2(t *testing.T) {
	first := d2()
	if first != "45973" {
		t.Fatalf("First Puzzle wrong with %v\n", first)
	}
	second := d2Part2()
	if second != "27CA4" {
		t.Fatalf("Second wrong with %v\n", second)
	}
}
