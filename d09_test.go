package main

import "testing"

// TestMain is test for d1
func TestD09(t *testing.T) {
	count1 := d09()
	if count1 != 97714 {
		t.Fatalf("Result1 %v is not ok \n", count1)
	}
	count2 := d09part2()
	if count2 != 10762972461 {
		t.Fatalf("Result2 %v is not ok \n", count2)
	}

}
