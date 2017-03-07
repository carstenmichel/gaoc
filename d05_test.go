package main

import "testing"

// TestMain is test for d1
func TestD05(t *testing.T) {
	first := d05()
	if first != "2414bc77" {
		t.Fatalf("Result %v is not ok \n", first)
	}
	second := d05Part2()
	if second != "437e60fc" {
		t.Fatalf("Wrong result in part 2 %v \n", second)
	}

}
