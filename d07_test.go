package main

import "testing"

// TestMain is test for d1
func TestD07(t *testing.T) {
	tls, ssl := d07()
	if tls != 110 {
		t.Fatalf("Result %v is not ok \n", tls)
	}
	if ssl != 242 {
		t.Fatalf("2nd Result %v is not ok \n", ssl)
	}

}
