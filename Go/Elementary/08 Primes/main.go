package main

import (
	"fmt"
)

// Calculates prime numbers indefinitely
func main() {
	i := 1
	fmt.Print("Prime numbers: ")
	for {
		if IsPrime(i) {
			fmt.Printf("%d, ", i)
		}

		i++
	}
}

func IsPrime(number int) bool {
	i := 1
	var isPrime bool = false

	for i < number/2 && !isPrime {
		if number%i == 0 {
			isPrime = true
		}
		i++
	}

	return isPrime
}
