package main

import "fmt"

func main() {
	var limit int = acceptInput()
	var sum int = calculateSum(limit)
	prettyPrintResult(limit, sum)
}

func acceptInput() int {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)
	return input
}

func calculateSum(limit int) int {
	var sum int

	for i := 1; i <= limit; i++ {
		sum += i
	}

	return sum
}

func prettyPrintResult(limit int, sum int) {
	for i := 1; i <= limit; i++ {
		if i == 1 {
			fmt.Printf("%d ", i)
		} else {
			fmt.Printf("+ %d ", i)
		}
	}

	fmt.Printf("= %d", sum)
}
