package main

import "fmt"

func fizzbuzz(n int) string {
	if n == 3 {
		return "Fizz"
	} else {
		return fmt.Sprintf("%d", n)
	}
}

func main() {}
