package main

import "testing"

func TestFizzbuzzReturns_1(t *testing.T) {
	result := fizzbuzz(1)
	expect := "1"

	if result != expect {
		t.Fail()
	}
}

func TestFizzbuzzReturns_Fizz(t *testing.T) {
	result := fizzbuzz(3)
	expect := "Fizz"

	if result != expect {
		t.Fail()
	}
}

func TestFizzbuzzReturns_Buzz(t *testing.T) {
	result := fizzbuzz(5)
	expect := "Buzz"

	if result != expect {
		t.Fail()
	}
}

func TestFizzbuzzReturns_Fizzbuzz(t *testing.T) {
	result := fizzbuzz(15)
	expect := "Fizzbuzz"

	if result != expect {
		t.Fail()
	}
}
