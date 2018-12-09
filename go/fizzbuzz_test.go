package main

import "testing"

func TestFizzbuzzReturns_1(t *testing.T) {
	result := fizzbuzz(1)
	expect := "1"

	if result != expect {
		t.Fail()
	}
}
