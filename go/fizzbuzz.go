package main

import "fmt"

func fizzbuzz(n int) string {
	if n == 15 {
		return "Fizzbuzz"
	} else if n == 5 {
		return "Buzz"
	} else if n == 3 {
		return "Fizz"
	} else {
		return fmt.Sprintf("%d", n)
	}
}

func main() {}
