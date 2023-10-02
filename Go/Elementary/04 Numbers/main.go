package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ApplyStrategy()
}

func ApplyStrategy() {
	var input int
	r := bufio.NewReader(os.Stdin)
	fmt.Println("Select an option:")
	fmt.Println("(1) Sum all numbers")
	fmt.Println("(2) Sum all numbers divisible by 3 or 5 only")
	fmt.Println("(3) Compute product of all numbers")
	var isValidInput bool = false

	for !isValidInput {
		fmt.Print("Which do you select (1, 2, 3)? ")
		_, err := fmt.Fscanf(r, "%d\n", &input)

		isValidInput = err == nil && (input >= 1 && input <= 3)

		if !isValidInput {
			fmt.Println("That is not a valid option.")
			r.ReadBytes(10)
		}
	}

	var limit int = AcceptInput()
	var sum int

	if input == 1 {
		sum = CalculateSum(limit)
		PrettyPrintResult(limit, sum)
	} else if input == 2 {
		sum = CalculateSumThreeOrFiveOnly(limit)
		PrettyPrintResultThreeOrFiveOnly(limit, sum)
	} else {
		var product uint64 = CalculateProduct(limit)
		PrettyPrintProduct(limit, product)
	}
}

func AcceptInput() int {
	var input int
	fmt.Print("Enter a number: ")
	fmt.Scan(&input)
	return input
}

func CalculateSum(limit int) int {
	var sum int

	for i := 1; i <= limit; i++ {
		sum += i
	}

	return sum
}

func CalculateSumThreeOrFiveOnly(limit int) int {
	var sum int

	for i := 1; i <= limit; i++ {
		if IsDivisibleByThreeOrFive(i) {
			sum += i
		}
	}

	return sum
}

func CalculateProduct(limit int) uint64 {
	var product uint64 = 1

	for i := uint64(1); i <= uint64(limit); i++ {
		product = product * i
	}

	return product
}

func PrettyPrintResult(limit int, sum int) {
	for i := 1; i <= limit; i++ {
		if i == 1 {
			fmt.Printf("%d ", i)
		} else {
			fmt.Printf("+ %d ", i)
		}
	}

	fmt.Printf("= %d", sum)
}

func PrettyPrintResultThreeOrFiveOnly(limit int, sum int) {
	for i := 1; i <= limit; i++ {
		if IsDivisibleByThreeOrFive(i) {
			if i == 3 {
				fmt.Printf("%d ", i)
			} else {
				fmt.Printf("+ %d ", i)
			}
		}
	}

	fmt.Printf("= %d", sum)
}

func PrettyPrintProduct(limit int, product uint64) {
	for i := 1; i <= limit; i++ {
		if i == 1 {
			fmt.Printf("%d ", i)
		} else {
			fmt.Printf("x %d ", i)
		}
	}

	fmt.Printf("= %d", product)
}

func IsDivisibleByThreeOrFive(number int) bool {
	return number%3 == 0 || number%5 == 0
}
