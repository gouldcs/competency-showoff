package main

import (
	"fmt"
)

func main() {
	var numberOfYears int
	var currentYear int
	fmt.Print("How many leap years would you like to see? ")
	fmt.Scanf("%d\n", &numberOfYears)

	fmt.Print("What year is it? ")
	fmt.Scanf("%d\n", &currentYear)

	var nextLeapYear int = findNextLeapYear(currentYear)
	printNextLeapYears(numberOfYears, nextLeapYear)

}

func findNextLeapYear(year int) int {
	if year%4 == 0 {
		return year + 4
	}

	nextLeapYear := year + 1

	for {
		if nextLeapYear%4 == 0 {
			return nextLeapYear
		}

		nextLeapYear++
	}
}

func printNextLeapYears(leapYearCount int, startYear int) {
	nextLeap := startYear

	fmt.Printf("The next %d leap years are: ", leapYearCount)
	for i := 0; i < leapYearCount; i++ {
		if i == 0 {
			fmt.Printf("%d", nextLeap)
		} else {
			fmt.Printf(", %d", nextLeap)
		}
		nextLeap += 4
	}
}
