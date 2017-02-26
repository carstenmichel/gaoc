package main

import "testing"

// TestMain is test for d1
func TestD1(t *testing.T) {
	short, long := d1Part1()
	if short != 163 {
		t.Fatal("Short path is wrong with %v\n", short)
	}
	if long != 279 {
		t.Fatal("Long path is wrong with %v\n", long)
	}
	t.Log("Day 1 test passed")
}
