package main

import "testing"

// TestMain is test for d1
func TestD06(t *testing.T) {
	first, second := d06()
	if first != "asvcbhvg" {
		t.Fatalf("Result %v is not ok \n", first)
	}
	if second != "odqnikqv" {
		t.Fatalf("2nd Result %v is not ok \n", first)
	}

}
