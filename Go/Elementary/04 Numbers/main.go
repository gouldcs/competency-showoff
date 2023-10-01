package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	applyStrategy()
}

func applyStrategy() {
	var input int
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Select an option:")
	fmt.Println("(1) Sum all numbers")
	fmt.Println("(2) Sum all numbers divisible by 3 or 5 only")
	var isValidInput bool = false

	for !isValidInput {
		fmt.Print("Which do you select (1, 2)? ")
		_, err := fmt.Fscanf(r, "%d\n", &input)

		isValidInput = err == nil && (input == 1 || input == 2)

		if !isValidInput {
			fmt.Println("That is not a valid option.")
			r.ReadBytes(10)
		}
	}

	var limit int = acceptInput()
	var sum int

	if input == 1 {
		sum = calculateSum(limit)
		prettyPrintResult(limit, sum)
	} else {
		sum = calculateSumThreeOrFiveOnly(limit)
		prettyPrintResultThreeOrFiveOnly(limit, sum)
	}
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

func calculateSumThreeOrFiveOnly(limit int) int {
	var sum int

	for i := 1; i <= limit; i++ {
		if isDivisibleByThreeOrFive(i) {
			sum += i
		}
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

func prettyPrintResultThreeOrFiveOnly(limit int, sum int) {
	for i := 1; i <= limit; i++ {
		if isDivisibleByThreeOrFive(i) {
			if i == 3 {
				fmt.Printf("%d ", i)
			} else {
				fmt.Printf("+ %d ", i)
			}
		}
	}

	fmt.Printf("= %d", sum)
}

func isDivisibleByThreeOrFive(number int) bool {
	return number%3 == 0 || number%5 == 0
}
