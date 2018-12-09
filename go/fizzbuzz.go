package main

import "fmt"

func fizzbuzz(n int) string {
	if n == 3 {
		return "Fizz"
	} else if n == 5 {
		return "Buzz"
	} else {
		return fmt.Sprintf("%d", n)
	}
}

func main() {}
