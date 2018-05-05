package main

import (
	"fmt"
	"math/rand"
)

func fizzbuzz(value int) string {
	if value%15 == 0 {
		return "fizz-buzz"
	}
	if value%3 == 0 {
		return "fizz"
	}
	if value%5 == 0 {
		return "buzz"
	}
	return fmt.Sprint(value)
}

func main() {
	fmt.Printf(fizzbuzz(rand.Intn(100)))
}
